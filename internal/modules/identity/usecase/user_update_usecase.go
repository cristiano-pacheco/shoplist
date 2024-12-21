package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type UserUpdateUseCase interface {
	Execute(ctx context.Context, input UserUpdateUseCaseInput) error
}

type UserUpdateUseCaseInput struct {
	UserID   uint64 `validate:"required"`
	Name     string `validate:"required,min=3,max=255"`
	Password string `validate:"required,min=8"`
}

type userUpdateUseCase struct {
	validate validator.Validate
	userRepo repository.UserRepository
	logger   logger.Logger
}

func NewUserUpdateUseCase(
	validate validator.Validate,
	userRepo repository.UserRepository,
	logger logger.Logger,
) UserUpdateUseCase {
	return &userUpdateUseCase{validate, userRepo, logger}
}

func (uc *userUpdateUseCase) Execute(ctx context.Context, input UserUpdateUseCaseInput) error {
	ctx, span := otel.Trace().StartSpan(ctx, "UserUpdateUseCase.Execute")
	defer span.End()

	err := uc.validate.Struct(input)
	if err != nil {
		return err
	}

	userModel, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		return err
	}

	userModel.Name = input.Name
	userModel.PasswordHash = input.Password

	err = uc.userRepo.Update(ctx, *userModel)
	if err != nil {
		message := "error updating user with id %d"
		uc.logger.Error(message, "error", err, "id", input.UserID)
		return err
	}

	return nil
}
