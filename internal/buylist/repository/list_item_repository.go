package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/buylist/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/errs"
)

type ListItemRepositoryI interface {
	Create(ctx context.Context, model model.ListItemModel) (*model.ListItemModel, error)
	Update(ctx context.Context, model model.ListItemModel) error
	Find(ctx context.Context, criteria FindListItemsCriteria) ([]*model.ListItemModel, error)
	Delete(ctx context.Context, criteria DeleteListItemCriteria) error
}

type listItemRepository struct {
	db *database.DB
}

func NewListItemRepository(db *database.DB) ListItemRepositoryI {
	return &listItemRepository{db}
}

func (r *listItemRepository) Create(ctx context.Context, model model.ListItemModel) (*model.ListItemModel, error) {
	result := r.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}

func (r *listItemRepository) Update(ctx context.Context, model model.ListItemModel) error {
	result := r.db.WithContext(ctx).Omit("updated_at").Save(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *listItemRepository) Find(ctx context.Context, criteria FindListItemsCriteria) ([]*model.ListItemModel, error) {
	var modelList []*model.ListItemModel
	r.db.WithContext(ctx).Where(criteria).Find(&modelList).Scan(&modelList)
	if len(modelList) == 0 {
		return nil, errs.ErrNotFound
	}
	return modelList, nil
}

func (r *listItemRepository) Delete(ctx context.Context, criteria DeleteListItemCriteria) error {
	result := r.db.WithContext(ctx).Delete(&model.ListItemModel{}, criteria)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
