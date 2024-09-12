package create_user_usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/service/hashservice"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
)

type UseCase struct {
	userRepo    repository.UserRepositoryI
	validate    validator.ValidateI
	hashService hashservice.HashServiceI
}

func New(
	userRepo repository.UserRepositoryI,
	validate validator.ValidateI,
	hashService hashservice.HashServiceI,

) *UseCase {
	return &UseCase{userRepo, validate, hashService}
}

func (uc *UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	err := uc.validate.Struct(input)
	if err != nil {
		return Output{}, err
	}

	ph, err := uc.hashService.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		return Output{}, err
	}

	userModel := model.UserModel{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: string(ph),
	}

	newUserModel, err := uc.userRepo.Create(ctx, userModel)
	if err != nil {
		return Output{}, err
	}

	output := Output{
		UserID: newUserModel.ID,
		Name:   newUserModel.Name,
		Email:  newUserModel.Email,
	}

	return output, nil
}
