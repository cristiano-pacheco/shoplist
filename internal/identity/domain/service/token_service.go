package service

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
)

type TokenService interface {
	Generate(ctx context.Context, user model.UserModel) (string, error)
}
