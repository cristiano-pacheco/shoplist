package repository

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
)

type AccountConfirmationRepository interface {
	Create(ctx context.Context, confirmation model.AccountConfirmationModel) (model.AccountConfirmationModel, error)
	FindByToken(ctx context.Context, token string) (model.AccountConfirmationModel, error)
	DeleteById(ctx context.Context, id uint64) error
	FindByUserID(ctx context.Context, userID uint64) (model.AccountConfirmationModel, error)
}
