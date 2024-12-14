package find_user

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
)

type FindUserUseCase struct {
	userRepo repository.UserRepository
	logger   logger.Logger
}

func New(
	userRepo repository.UserRepository,
	logger logger.Logger,
) *FindUserUseCase {
	return &FindUserUseCase{userRepo, logger}
}

func (uc *FindUserUseCase) Execute(ctx context.Context, input Input) (Output, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "FindUserUseCase.Execute")
	defer span.End()

	userModel, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		message := "[find_user] error finding user by id %d"
		uc.logger.Error(message, "error", err, "id", input.UserID)
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
