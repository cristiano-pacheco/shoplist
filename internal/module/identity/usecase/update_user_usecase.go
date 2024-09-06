package usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
	"github.com/go-playground/validator/v10"
)

type UpdateUserUseCase struct {
	validate *validator.Validate
	userRepo repository.UserRepositoryI
}

func NewUpdateUserUseCase(
	validate *validator.Validate,
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
