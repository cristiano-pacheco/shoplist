package repository

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/shared/database"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
)

type ListRepositoryI interface {
	Create(ctx context.Context, model model.ListModel) (*model.ListModel, error)
	Update(ctx context.Context, model model.ListModel) error
	Find(ctx context.Context, criteria FindListsCriteria) ([]*model.ListModel, error)
	Delete(ctx context.Context, criteria DeleteListCriteria) error
}

type listRepository struct {
	db *database.ShoplistDB
}

func NewListRepository(db *database.ShoplistDB) ListRepositoryI {
	return &listRepository{db}
}

func (r *listRepository) Create(ctx context.Context, model model.ListModel) (*model.ListModel, error) {
	result := r.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}

func (r *listRepository) Update(ctx context.Context, model model.ListModel) error {
	result := r.db.WithContext(ctx).Omit("updated_at").Save(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *listRepository) Find(ctx context.Context, criteria FindListsCriteria) ([]*model.ListModel, error) {
	var modelList []*model.ListModel
	r.db.WithContext(ctx).Where(criteria).Find(&modelList).Scan(&modelList)
	if len(modelList) == 0 {
		return nil, errs.ErrNotFound
	}
	return modelList, nil
}

func (r *listRepository) Delete(ctx context.Context, criteria DeleteListCriteria) error {
	result := r.db.WithContext(ctx).Delete(&model.ListModel{}, criteria)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
