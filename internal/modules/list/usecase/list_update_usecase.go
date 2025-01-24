package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/errs"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type ListUpdateUseCase struct {
	listRepository repository.ListRepository
	validate       validator.Validate
}

type ListUpdateInput struct {
	UserID uint64 `validate:"required"`
	ListID uint64 `validate:"required"`
	Name   string `validate:"required,min=1,max=255"`
}

func NewListUpdateUseCase(
	listRepository repository.ListRepository,
	validate validator.Validate,
) *ListUpdateUseCase {
	return &ListUpdateUseCase{listRepository, validate}
}

func (uc *ListUpdateUseCase) Execute(ctx context.Context, input ListUpdateInput) error {
	err := uc.validate.Struct(input)
	if err != nil {
		return err
	}

	model, err := uc.listRepository.FindByIDAndUserID(ctx, input.ListID, input.UserID)
	if err != nil {
		return err
	}

	if model.UserID != input.UserID {
		return errs.ErrResourceDoesNotBelongToUser
	}

	model.Name = input.Name
	err = uc.listRepository.Update(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
