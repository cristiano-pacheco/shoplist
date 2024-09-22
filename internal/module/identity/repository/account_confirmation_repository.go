package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
)

type AccountConfirmationRepositoryI interface {
	Create(ctx context.Context, m model.AccountConfirmationModel) error
	FindByUserID(ctx context.Context, userID uint64) (model.AccountConfirmationModel, error)
	Delete(ctx context.Context, m model.AccountConfirmationModel) error
}

type accountConfirmationRepository struct {
	db *database.DB
}

func NewAccountConfirmationRepository(db *database.DB) AccountConfirmationRepositoryI {
	return &accountConfirmationRepository{db}
}

func (r *accountConfirmationRepository) Create(ctx context.Context, m model.AccountConfirmationModel) error {
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *accountConfirmationRepository) FindByUserID(
	ctx context.Context,
	userID uint64,
) (model.AccountConfirmationModel, error) {
	var m model.AccountConfirmationModel
	return m, r.db.WithContext(ctx).Where("user_id = ?", userID).First(&m).Error
}

func (r *accountConfirmationRepository) Delete(ctx context.Context, m model.AccountConfirmationModel) error {
	return r.db.WithContext(ctx).Delete(&m).Error
}
