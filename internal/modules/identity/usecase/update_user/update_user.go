package update_user

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type UpdateUserUseCase struct {
	validate validator.ValidateI
	userRepo repository.UserRepositoryI
	logger   logger.LoggerI
}

func New(
	validate validator.ValidateI,
	userRepo repository.UserRepositoryI,
	logger logger.LoggerI,
) *UpdateUserUseCase {
	return &UpdateUserUseCase{validate, userRepo, logger}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, input Input) error {
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
		message := "[update_user_usecase] error updating user"
		uc.logger.Error(message, "error", err)
		return err
	}

	return nil
}
