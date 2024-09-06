package usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
)

type FindUserByIDUseCase struct {
	userRepo repository.UserRepositoryI
}

func NewFindUserByIDUseCase(
	userRepo repository.UserRepositoryI,
) *FindUserByIDUseCase {
	return &FindUserByIDUseCase{userRepo}
}

func (uc *FindUserByIDUseCase) Execute(
	ctx context.Context,
	input dto.FindUserByIDInput,
) (dto.FindUserByIDOutput, error) {
	userModel, err := uc.userRepo.FindOneByID(ctx, input.UserID)
	if err != nil {
		return dto.FindUserByIDOutput{}, err
	}

	output := dto.FindUserByIDOutput{
		UserID:   userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.PasswordHash,
	}

	return output, nil
}
