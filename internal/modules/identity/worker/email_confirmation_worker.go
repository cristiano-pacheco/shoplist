package worker

import (
	"context"
)

type EmailConfirmationWorker interface {
	Start()
	Enqueue(ctx context.Context, userID int64) error
}
