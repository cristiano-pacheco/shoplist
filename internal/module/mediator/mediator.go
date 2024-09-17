package mediator

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/find_activated_user_by_id_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mediator"
)

type Mediator struct {
	findActivatedUserByIdUseCase *find_activated_user_by_id_usecase.UseCase
}

func New(
	findActivatedUserByIdUseCase *find_activated_user_by_id_usecase.UseCase,
) mediator.MediatorI {
	return &Mediator{
		findActivatedUserByIdUseCase,
	}
}

func (m *Mediator) IsUserActivated(ctx context.Context, userID uint64) bool {
	input := find_activated_user_by_id_usecase.Input{UserID: userID}
	return m.findActivatedUserByIdUseCase.Execute(ctx, input).IsUserActivated
}
