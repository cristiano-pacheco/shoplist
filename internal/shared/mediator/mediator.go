package mediator

import "context"

type MediatorI interface {
	IsUserActivated(ctx context.Context, userID uint64) bool
}
