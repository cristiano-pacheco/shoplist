package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/list/errs"
	"github.com/cristiano-pacheco/shoplist/internal/list/repository"
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

	model, err := uc.listRepository.FindByIDAndUserID(ctx, input.ListID, input.UserID)
	if err != nil {
		return err
	}

	if model.UserID != input.UserID {
		return errs.ErrResourceDoesNotBelongToUser
	}

	err = uc.listRepository.Delete(ctx, repository.DeleteListCriteria{
		ID:     input.ListID,
		UserID: input.UserID,
	})

	if err != nil {
		return err
	}

	return nil
}
