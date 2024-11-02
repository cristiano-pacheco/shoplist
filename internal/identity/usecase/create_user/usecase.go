package create_user_usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/identity/model"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/repository"
	hash_service "github.com/cristiano-pacheco/go-modulith/internal/identity/service/hash"
	send_account_confirmation_email_service "github.com/cristiano-pacheco/go-modulith/internal/identity/service/send_account_confirmation_email"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/logger"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
)

type UseCase struct {
	userRepo                            repository.UserRepositoryI
	validate                            validator.ValidateI
	hashService                         hash_service.ServiceI
	sendAccountConfirmationEmailService send_account_confirmation_email_service.ServiceI
	logger                              logger.LoggerI
}

func New(
	userRepo repository.UserRepositoryI,
	validate validator.ValidateI,
	hashService hash_service.ServiceI,
	sendAccountConfirmationEmailService send_account_confirmation_email_service.ServiceI,
	logger logger.LoggerI,
) *UseCase {
	return &UseCase{userRepo, validate, hashService, sendAccountConfirmationEmailService, logger}
}

func (uc *UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "create_user_usecase.Execute")
	defer span.End()

	err := uc.validate.Struct(input)
	if err != nil {
		return Output{}, err
	}

	ph, err := uc.hashService.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		message := "[create_user_usecase] error generating password hash"
		uc.logger.Error(message, "error", err)
		return Output{}, err
	}

	userModel := model.UserModel{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: string(ph),
	}

	newUserModel, err := uc.userRepo.Create(ctx, userModel)
	if err != nil {
		message := "[create_user_usecase] error creating user"
		uc.logger.Error(message, "error", err)
		return Output{}, err
	}

	err = uc.sendAccountConfirmationEmailService.Execute(ctx, *newUserModel)
	if err != nil {
		message := "[create_user_usecase] error sending account confirmation email"
		uc.logger.Error(message, "error", err)
		return Output{}, err
	}

	output := Output{
		UserID: newUserModel.ID,
		Name:   newUserModel.Name,
		Email:  newUserModel.Email,
	}

	return output, nil
}
