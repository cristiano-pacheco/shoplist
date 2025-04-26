package repository

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
)

type LoginTokenRepository interface {
	Create(ctx context.Context, loginToken model.LoginTokenModel) (model.LoginTokenModel, error)
	Update(ctx context.Context, loginToken model.LoginTokenModel) error
	Delete(ctx context.Context, id uint64) error
	FindByToken(ctx context.Context, token string) (model.LoginTokenModel, error)
}
