package mediator

import (
	"context"
	"errors"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/find_activated_user_by_id_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mediator"
)

type Mediator struct {
}

func New(
	findActivatedUserByIdUseCase *find_activated_user_by_id_usecase.UseCase,
) mediator.MediatorI {
	mediator := mediator.New()

	callback := func(ctx context.Context, input any) (any, error) {
		userID, ok := input.(uint64)
		if !ok {
			return nil, errors.New("invalid input")
		}
		uInput := find_activated_user_by_id_usecase.Input{UserID: userID}
		output := findActivatedUserByIdUseCase.Execute(ctx, uInput)
		return output.IsUserActivated, nil
	}

	mediator.Register("is_user_activated", callback)
	return mediator
}
