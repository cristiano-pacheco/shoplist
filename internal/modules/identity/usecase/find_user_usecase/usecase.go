package find_user_usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/logger"
)

type UseCase struct {
	userRepo repository.UserRepositoryI
	logger   logger.LoggerI
}

func New(
	userRepo repository.UserRepositoryI,
	logger logger.LoggerI,
) *UseCase {
	return &UseCase{userRepo, logger}
}

func (uc *UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	userModel, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		message := "[find_user_by_id_usecase] error finding user by id"
		uc.logger.Error(message, "error", err)
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
