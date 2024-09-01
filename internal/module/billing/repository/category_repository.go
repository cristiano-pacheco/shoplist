package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/module/database"
)

type CategoryRepositoryI interface {
	Create(ctx context.Context, model model.CategoryModel) (*model.CategoryModel, error)
	Update(ctx context.Context, model model.CategoryModel) (*model.CategoryModel, error)
	Delete(ctx context.Context, model model.CategoryModel) error
	FindOneByID(ctx context.Context, ID uint64) (*model.CategoryModel, error)
	FindList(ctx context.Context, filters CategorySearchFilters) []*model.CategoryModel
}

type CategorySearchFilters struct {
	UserID *uint64
}

type categoryRepository struct {
	db *database.DB
}

func NewCategoryRepository(db *database.DB) CategoryRepositoryI {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(
	ctx context.Context,
	model model.CategoryModel,
) (*model.CategoryModel, error) {
	return nil, nil
}

func (r *categoryRepository) Update(
	ctx context.Context,
	model model.CategoryModel,
) (*model.CategoryModel, error) {
	return nil, nil
}

func (r *categoryRepository) Delete(
	ctx context.Context,
	model model.CategoryModel,
) error {
	return nil
}

func (r *categoryRepository) FindOneByID(
	ctx context.Context,
	ID uint64,
) (*model.CategoryModel, error) {
	return nil, nil
}

func (r *categoryRepository) FindList(
	ctx context.Context,
	filters CategorySearchFilters,
) []*model.CategoryModel {
	return nil
}
