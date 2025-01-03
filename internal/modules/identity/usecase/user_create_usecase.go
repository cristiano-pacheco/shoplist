package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type UserCreateUseCase interface {
	Execute(ctx context.Context, input UserCreateUseCaseInput) (UserCreateUseCaseOutput, error)
}

type UserCreateUseCaseInput struct {
	Name     string `validate:"required,min=3,max=255"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type UserCreateUseCaseOutput struct {
	Name   string
	Email  string
	UserID uint64
}

type userCreateUseCase struct {
	emailConfirmationService service.EmailConfirmationService
	hashService              service.HashService
	userRepo                 repository.UserRepository
	validate                 validator.Validate
	logger                   logger.Logger
}

func NewUserCreateUseCase(
	emailConfirmationService service.EmailConfirmationService,
	hashService service.HashService,
	userRepo repository.UserRepository,
	validate validator.Validate,
	logger logger.Logger,
) UserCreateUseCase {
	return &userCreateUseCase{
		emailConfirmationService,
		hashService,
		userRepo,
		validate,
		logger,
	}
}

func (uc *userCreateUseCase) Execute(ctx context.Context, input UserCreateUseCaseInput) (UserCreateUseCaseOutput, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserCreateUseCase.Execute")
	defer span.End()

	output := UserCreateUseCaseOutput{}

	err := uc.validate.Struct(input)
	if err != nil {
		return output, err
	}

	ph, err := uc.hashService.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		message := "error generating password hash"
		uc.logger.Error(message, "error", err)
		return output, err
	}

	userModel := model.UserModel{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: string(ph),
	}

	newUserModel, err := uc.userRepo.Create(ctx, userModel)
	if err != nil {
		message := "error creating user"
		uc.logger.Error(message, "error", err)
		return output, err
	}

	err = uc.emailConfirmationService.Send(ctx, newUserModel.ID)
	if err != nil {
		message := "error sending account confirmation email"
		uc.logger.Error(message, "error", err)
		return output, err
	}

	output = UserCreateUseCaseOutput{
		UserID: newUserModel.ID,
		Name:   newUserModel.Name,
		Email:  newUserModel.Email,
	}

	return output, nil
}
