package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/persistence"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
)

type CategoryRepositoryI interface {
	Create(ctx context.Context, model persistence.CategoryModel) (*persistence.CategoryModel, error)
	Update(ctx context.Context, model persistence.CategoryModel) (*persistence.CategoryModel, error)
	Delete(ctx context.Context, model persistence.CategoryModel) error
	FindOneByID(ctx context.Context, ID uint64) (*persistence.CategoryModel, error)
	FindList(ctx context.Context, filters CategorySearchFilters) []*persistence.CategoryModel
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
	model persistence.CategoryModel,
) (*persistence.CategoryModel, error) {
	return nil, nil
}

func (r *categoryRepository) Update(
	ctx context.Context,
	model persistence.CategoryModel,
) (*persistence.CategoryModel, error) {
	return nil, nil
}

func (r *categoryRepository) Delete(
	ctx context.Context,
	model persistence.CategoryModel,
) error {
	return nil
}

func (r *categoryRepository) FindOneByID(
	ctx context.Context,
	ID uint64,
) (*persistence.CategoryModel, error) {
	return nil, nil
}

func (r *categoryRepository) FindList(
	ctx context.Context,
	filters CategorySearchFilters,
) []*persistence.CategoryModel {
	return nil
}
