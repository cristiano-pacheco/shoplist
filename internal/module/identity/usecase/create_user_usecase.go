package usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
)

type CreateUserUseCase struct {
	userRepo repository.UserRepositoryI
	validate validator.ValidateI
}

func NewCreateUserUseCaseUseCase(
	userRepo repository.UserRepositoryI,
	validate validator.ValidateI,
) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo, validate}
}

func (uc *CreateUserUseCase) Execute(
	ctx context.Context,
	input dto.CreateUserInput,
) (dto.CreateUserOutput, error) {
	err := uc.validate.Struct(input)
	if err != nil {
		return dto.CreateUserOutput{}, err
	}

	userModel := model.UserModel{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: input.Password,
	}

	newUserModel, err := uc.userRepo.Create(ctx, userModel)
	if err != nil {
		return dto.CreateUserOutput{}, err
	}

	output := dto.CreateUserOutput{
		UserID: newUserModel.ID,
		Name:   newUserModel.Name,
		Email:  newUserModel.Email,
	}

	return output, nil
}
