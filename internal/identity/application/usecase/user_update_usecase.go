package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/service"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/logger"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/otel"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/validator"
)

type UserUpdateUseCase interface {
	Execute(ctx context.Context, input UserUpdateInput) error
}

type UserUpdateInput struct {
	UserID   uint64 `validate:"required"`
	Name     string `validate:"required,min=3,max=255"`
	Password string `validate:"required,min=8"`
}

type userUpdateUseCase struct {
	validate    validator.Validate
	userRepo    repository.UserRepository
	logger      logger.Logger
	hashService service.HashService
}

func NewUserUpdateUseCase(
	validate validator.Validate,
	userRepo repository.UserRepository,
	logger logger.Logger,
	hashService service.HashService,
) UserUpdateUseCase {
	return &userUpdateUseCase{validate, userRepo, logger, hashService}
}

func (uc *userUpdateUseCase) Execute(ctx context.Context, input UserUpdateInput) error {
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

	ph, err := uc.hashService.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		message := "error generating password hash"
		uc.logger.Error(message, "error", err)
		return err
	}

	updatedUserModel, err := model.RestoreUserModel(
		userModel.ID(),
		input.Name,
		userModel.Email(),
		string(ph),
		userModel.IsActivated(),
		userModel.RpToken(),
		userModel.CreatedAt(),
		userModel.UpdatedAt(),
	)
	if err != nil {
		return err
	}

	err = uc.userRepo.Update(ctx, updatedUserModel)
	if err != nil {
		message := "error updating user with id %d"
		uc.logger.Error(message, "error", err, "id", input.UserID)
		return err
	}

	return nil
}
