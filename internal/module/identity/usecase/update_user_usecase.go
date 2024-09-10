package usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
)

type UpdateUserUseCase struct {
	validate validator.ValidateI
	userRepo repository.UserRepositoryI
}

func NewUpdateUserUseCase(
	validate validator.ValidateI,
	userRepo repository.UserRepositoryI,
) *UpdateUserUseCase {
	return &UpdateUserUseCase{validate, userRepo}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, input dto.UpdateUserInput) error {
	err := uc.validate.Struct(input)
	if err != nil {
		return err
	}

	userModel, err := uc.userRepo.FindOneByID(ctx, input.UserID)
	if err != nil {
		return err
	}

	userModel.Name = input.Name
	userModel.PasswordHash = input.Password

	err = uc.userRepo.Update(ctx, *userModel)
	if err != nil {
		return err
	}

	return nil
}
