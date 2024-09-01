package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/module/database"
)

type BillingRepositoryI interface {
	Create(ctx context.Context, model model.BillingModel) (*model.BillingModel, error)
	Update(ctx context.Context, model model.BillingModel) (*model.BillingModel, error)
	Delete(ctx context.Context, model model.BillingModel) error
	FindOneByID(ctx context.Context, ID uint64) (*model.BillingModel, error)
	FindList(ctx context.Context, filters BillingSearchFilters) []*model.BillingModel
}

type BillingSearchFilters struct {
	UserID     *uint64
	CategoryID *uint64
}

type billingRepository struct {
	db *database.DB
}

func NewBillingRepository(db *database.DB) BillingRepositoryI {
	return &billingRepository{db}
}

func (r *billingRepository) Create(
	ctx context.Context,
	model model.BillingModel,
) (*model.BillingModel, error) {
	return nil, nil
}

func (r *billingRepository) Update(
	ctx context.Context,
	model model.BillingModel,
) (*model.BillingModel, error) {
	return nil, nil
}

func (r *billingRepository) Delete(ctx context.Context, model model.BillingModel) error {
	return nil
}

func (r *billingRepository) FindOneByID(ctx context.Context, ID uint64) (*model.BillingModel, error) {
	return nil, nil
}

func (r *billingRepository) FindList(ctx context.Context, filters BillingSearchFilters) []*model.BillingModel {
	return nil
}
