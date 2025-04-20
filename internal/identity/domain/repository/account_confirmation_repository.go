package repository

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
)

type AccountConfirmationRepository interface {
	Create(ctx context.Context, confirmation model.AccountConfirmationModel) (model.AccountConfirmationModel, error)
	FindByToken(ctx context.Context, token string) (model.AccountConfirmationModel, error)
	DeleteById(ctx context.Context, id uint) error
	FindByUserID(ctx context.Context, userID uint) (model.AccountConfirmationModel, error)
}
