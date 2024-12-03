package create_user

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/telemetry"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type CreateUserUseCase struct {
	sendAccountConfirmationEmailService service.SendAccountConfirmationEmailServiceI
	hashService                         service.HashServiceI
	userRepo                            repository.UserRepositoryI
	validate                            validator.ValidateI
	logger                              logger.LoggerI
}

func New(
	sendAccountConfirmationEmailService service.SendAccountConfirmationEmailServiceI,
	hashService service.HashServiceI,
	userRepo repository.UserRepositoryI,
	validate validator.ValidateI,
	logger logger.LoggerI,
) *CreateUserUseCase {
	return &CreateUserUseCase{sendAccountConfirmationEmailService, hashService, userRepo, validate, logger}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input Input) (Output, error) {
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
