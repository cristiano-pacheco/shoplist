package repository

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/shared/database"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
)

type ListRepository interface {
	Create(ctx context.Context, model model.ListModel) (model.ListModel, error)
	Update(ctx context.Context, model model.ListModel) error
	FindByIDAndUserID(ctx context.Context, id uint64, userID uint64) (model.ListModel, error)
	FindByUserID(ctx context.Context, userID uint64) ([]model.ListModel, error)
	Delete(ctx context.Context, criteria DeleteListCriteria) error
}

type listRepository struct {
	db *database.ShoplistDB
}

func NewListRepository(db *database.ShoplistDB) ListRepository {
	return &listRepository{db}
}

func (r *listRepository) Create(ctx context.Context, model model.ListModel) (model.ListModel, error) {
	result := r.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return model, result.Error
	}
	return model, nil
}

func (r *listRepository) Update(ctx context.Context, model model.ListModel) error {
	result := r.db.WithContext(ctx).Omit("updated_at").Save(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *listRepository) FindByUserID(ctx context.Context, userID uint64) ([]model.ListModel, error) {
	var modelList []model.ListModel
	result := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&modelList)
	if result.Error != nil {
		return nil, result.Error
	}
	return modelList, nil
}

func (r *listRepository) FindByIDAndUserID(ctx context.Context, id uint64, userID uint64) (model.ListModel, error) {
	var model model.ListModel
	result := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&model)

	if result.Error != nil {
		return model, result.Error
	}
	return model, nil
}

func (r *listRepository) Delete(ctx context.Context, criteria DeleteListCriteria) error {
	result := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", criteria.ID, criteria.UserID).
		Delete(&model.ListModel{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}
