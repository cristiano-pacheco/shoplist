package find_activated_user_by_id_usecase

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

func (uc *UseCase) Execute(ctx context.Context, input Input) Output {
	isUserActivated := uc.userRepo.IsUsedActivated(ctx, input.UserID)

	output := Output{
		IsUserActivated: isUserActivated,
	}

	return output
}
