package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
)

type UserFindUseCase interface {
	Execute(ctx context.Context, input UserFindUseCaseInput) (UserFindUseCaseOutput, error)
}

type UserFindUseCaseInput struct {
	UserID uint64 `validate:"required,number"`
}

type UserFindUseCaseOutput struct {
	UserID   uint64
	Name     string
	Email    string
	Password string
}

type userFindUseCase struct {
	userRepo repository.UserRepository
	logger   logger.Logger
}

func NewUserFindUseCase(
	userRepo repository.UserRepository,
	logger logger.Logger,
) UserFindUseCase {
	return &userFindUseCase{userRepo, logger}
}

func (uc *userFindUseCase) Execute(ctx context.Context, input UserFindUseCaseInput) (UserFindUseCaseOutput, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserFindUseCase.Execute")
	defer span.End()

	userModel, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		message := "error finding user by id %d"
		uc.logger.Error(message, "error", err, "id", input.UserID)
		return UserFindUseCaseOutput{}, err
	}

	output := UserFindUseCaseOutput{
		UserID:   userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.PasswordHash,
	}

	return output, nil
}
