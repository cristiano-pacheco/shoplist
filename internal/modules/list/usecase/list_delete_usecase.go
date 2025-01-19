package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/errs"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type ListDeleteUseCase struct {
	listRepository repository.ListRepository
	validate       validator.Validate
}

type ListDeleteInput struct {
	UserID uint64 `validate:"required"`
	ListID uint64 `validate:"required"`
}

func NewListDeleteUseCase(
	listRepository repository.ListRepository,
	validate validator.Validate,
) *ListDeleteUseCase {
	return &ListDeleteUseCase{listRepository, validate}
}

func (uc *ListDeleteUseCase) Execute(ctx context.Context, input ListDeleteInput) error {
	err := uc.validate.Struct(input)
	if err != nil {
		return err
	}

	criteria := repository.DeleteListCriteria{
		UserID: input.UserID,
		ListID: input.ListID,
	}

	model, err := uc.listRepository.FindByID(ctx, input.ListID)
	if err != nil {
		return err
	}

	if model.UserID != input.UserID {
		return errs.ErrResourceDoesNotBelongToUser
	}

	err = uc.listRepository.Delete(ctx, criteria)
	if err != nil {
		return err
	}

	return nil
}
