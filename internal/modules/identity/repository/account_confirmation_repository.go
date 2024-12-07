package repository

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/model"
	"github.com/cristiano-pacheco/shoplist/internal/shared/database"
	"github.com/cristiano-pacheco/shoplist/internal/shared/telemetry"
)

type AccountConfirmationRepository interface {
	Create(ctx context.Context, m model.AccountConfirmationModel) error
	FindByUserID(ctx context.Context, userID uint64) (model.AccountConfirmationModel, error)
	Delete(ctx context.Context, m model.AccountConfirmationModel) error
}

type accountConfirmationRepository struct {
	db *database.ShoplistDB
}

func NewAccountConfirmationRepository(db *database.ShoplistDB) AccountConfirmationRepository {
	return &accountConfirmationRepository{db}
}

func (r *accountConfirmationRepository) Create(ctx context.Context, m model.AccountConfirmationModel) error {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "account_confirmation_repository.create")
	defer span.End()
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *accountConfirmationRepository) FindByUserID(
	ctx context.Context,
	userID uint64,
) (model.AccountConfirmationModel, error) {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "account_confirmation_repository.find_by_user_id")
	defer span.End()
	var m model.AccountConfirmationModel
	return m, r.db.WithContext(ctx).Where("user_id = ?", userID).First(&m).Error
}

func (r *accountConfirmationRepository) Delete(ctx context.Context, m model.AccountConfirmationModel) error {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "account_confirmation_repository.delete")
	defer span.End()
	return r.db.WithContext(ctx).Where("user_id = ?", m.UserID).Delete(&m).Error
}
