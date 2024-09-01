package usecase

import (
	"github.com/cristiano-pacheco/go-modulith/internal/modules/billing/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/modules/billing/usecase/data"
	"github.com/go-playground/validator/v10"
)

type CreateBillingUseCase struct {
	billingRepo repository.BillingRepositoryI
	validate    *validator.Validate
}

func NewCreateBillingUseCase(
	billingRepo repository.BillingRepositoryI,
	validate *validator.Validate,
) *CreateBillingUseCase {
	return &CreateBillingUseCase{billingRepo, validate}
}

func (uc *CreateBillingUseCase) Execute(input data.CreateBillingInput) (data.CreateBillingOutput, error) {
	err := uc.validate.Struct(input)
	if err != nil {
		//validationErrors := err.(validator.ValidationErrors)
		return data.CreateBillingOutput{}, err
	}

	return data.CreateBillingOutput{}, nil
}
