package usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/service/hashservice"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
)

type CreateUserUseCase struct {
	userRepo    repository.UserRepositoryI
	validate    validator.ValidateI
	hashService hashservice.HashServiceI
}

func NewCreateUserUseCaseUseCase(
	userRepo repository.UserRepositoryI,
	validate validator.ValidateI,
	hashService hashservice.HashServiceI,

) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo, validate, hashService}
}

func (uc *CreateUserUseCase) Execute(
	ctx context.Context,
	input dto.CreateUserInput,
) (dto.CreateUserOutput, error) {
	err := uc.validate.Struct(input)
	if err != nil {
		return dto.CreateUserOutput{}, err
	}

	ph, err := uc.hashService.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		return dto.CreateUserOutput{}, err
	}

	userModel := model.UserModel{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: string(ph),
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
