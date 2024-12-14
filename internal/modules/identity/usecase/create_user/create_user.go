package create_user

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
	"go.opentelemetry.io/otel/trace"
)

type CreateUserUseCase struct {
	emailConfirmationService service.EmailConfirmationService
	hashService              service.HashService
	userRepo                 repository.UserRepository
	validate                 validator.Validate
	logger                   logger.Logger
}

func New(
	emailConfirmationService service.EmailConfirmationService,
	hashService service.HashService,
	userRepo repository.UserRepository,
	validate validator.Validate,
	logger logger.Logger,
) *CreateUserUseCase {
	return &CreateUserUseCase{emailConfirmationService, hashService, userRepo, validate, logger}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input Input) (Output, error) {
	span := trace.SpanFromContext(ctx)
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

	err = uc.emailConfirmationService.Send(ctx, newUserModel.ID)
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
