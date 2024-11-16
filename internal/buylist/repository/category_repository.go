package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/buylist/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/errs"
)

type CategoryRepositoryI interface {
	Create(ctx context.Context, model model.CategoryModel) (*model.CategoryModel, error)
	Update(ctx context.Context, model model.CategoryModel) error
	Find(ctx context.Context, criteria FindCategoriesCriteria) ([]*model.CategoryModel, error)
	Delete(ctx context.Context, criteria DeleteCategoryCriteria) error
}

type categoryRepository struct {
	db *database.DB
}

func NewcategoryRepository(db *database.DB) CategoryRepositoryI {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(ctx context.Context, model model.CategoryModel) (*model.CategoryModel, error) {
	result := r.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}

func (r *categoryRepository) Update(ctx context.Context, model model.CategoryModel) error {
	result := r.db.WithContext(ctx).Omit("updated_at").Save(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *categoryRepository) Find(ctx context.Context, criteria FindCategoriesCriteria) ([]*model.CategoryModel, error) {
	var modelList []*model.CategoryModel
	r.db.WithContext(ctx).Where(criteria).Find(&modelList).Scan(&modelList)
	if len(modelList) == 0 {
		return nil, errs.ErrNotFound
	}
	return modelList, nil
}

func (r *categoryRepository) Delete(ctx context.Context, criteria DeleteCategoryCriteria) error {
	result := r.db.WithContext(ctx).Delete(&model.CategoryModel{}, criteria)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
