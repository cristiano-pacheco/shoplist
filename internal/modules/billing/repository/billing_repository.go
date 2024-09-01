package repository

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/persistence"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/database"
)

type BillingRepositoryI interface {
	Create(ctx context.Context, model persistence.BillingModel) (*persistence.BillingModel, error)
	Update(ctx context.Context, model persistence.BillingModel) (*persistence.BillingModel, error)
	Delete(ctx context.Context, model persistence.BillingModel) error
	FindOneByID(ctx context.Context, ID uint64) (*persistence.BillingModel, error)
	FindList(ctx context.Context, filters BillingSearchFilters) []*persistence.BillingModel
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
	model persistence.BillingModel,
) (*persistence.BillingModel, error) {
	return nil, nil
}

func (r *billingRepository) Update(
	ctx context.Context,
	model persistence.BillingModel,
) (*persistence.BillingModel, error) {
	return nil, nil
}

func (r *billingRepository) Delete(ctx context.Context, model persistence.BillingModel) error {
	return nil
}

func (r *billingRepository) FindOneByID(ctx context.Context, ID uint64) (*persistence.BillingModel, error) {
	return nil, nil
}

func (r *billingRepository) FindList(ctx context.Context, filters BillingSearchFilters) []*persistence.BillingModel {
	return nil
}
