package update_user_usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
)

type UseCase struct {
	validate validator.ValidateI
	userRepo repository.UserRepositoryI
}

func New(
	validate validator.ValidateI,
	userRepo repository.UserRepositoryI,
) *UseCase {
	return &UseCase{validate, userRepo}
}

func (uc *UseCase) Execute(ctx context.Context, input Input) error {
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
		return err
	}

	return nil
}
