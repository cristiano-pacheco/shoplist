package find_user_by_id_usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
)

type UseCase struct {
	userRepo repository.UserRepositoryI
}

func New(
	userRepo repository.UserRepositoryI,
) *UseCase {
	return &UseCase{userRepo}
}

func (uc *UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	userModel, err := uc.userRepo.FindOneByID(ctx, input.UserID)
	if err != nil {
		return Output{}, err
	}

	output := Output{
		UserID:   userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.PasswordHash,
	}

	return output, nil
}
