package usecase

import (
	"context"

	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/validator"
)

type CreateBillingUseCase struct {
	billingRepo repository.BillingRepositoryI
	validate    validator.ValidateI
}

func NewCreateBillingUseCase(
	billingRepo repository.BillingRepositoryI,
	validate validator.ValidateI,
) *CreateBillingUseCase {
	return &CreateBillingUseCase{billingRepo, validate}
}

func (uc *CreateBillingUseCase) Execute(
	ctx context.Context,
	input dto.CreateBillingInput,
) (dto.CreateBillingOutput, error) {
	err := uc.validate.Struct(input)
	if err != nil {
		return dto.CreateBillingOutput{}, err
	}

	return dto.CreateBillingOutput{}, nil
}
