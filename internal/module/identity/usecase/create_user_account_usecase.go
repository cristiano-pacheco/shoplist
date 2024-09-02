package usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
	"github.com/go-playground/validator/v10"
)

type CreateUserAccountUseCase struct {
	billingRepo repository.UserRepositoryI
	validate    *validator.Validate
}

func NewCreateUserAccountUseCase(
	userRepo repository.UserRepositoryI,
	validate *validator.Validate,
) *CreateUserAccountUseCase {
	return &CreateUserAccountUseCase{userRepo, validate}
}

func (uc *CreateUserAccountUseCase) Execute(
	ctx context.Context,
	input dto.CreateUserAccountInput,
) (dto.CreateUserAccountOutput, error) {
	err := uc.validate.Struct(input)
	if err != nil {
		//validationErrors := err.(validator.ValidationErrors)
		return dto.CreateUserAccountOutput{}, err
	}

	return dto.CreateUserAccountOutput{}, nil
}
