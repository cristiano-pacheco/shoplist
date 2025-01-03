package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type CategoryDeleteUseCase struct {
	categoryRepository repository.CategoryRepository
	validate           validator.Validate
}

type CategoryDeleteInput struct {
	UserID     uint64 `validate:"required"`
	CategoryID uint64 `validate:"required"`
}

func NewCategoryDeleteUseCase(
	categoryRepository repository.CategoryRepository,
	validate validator.Validate,
) *CategoryDeleteUseCase {
	return &CategoryDeleteUseCase{categoryRepository, validate}
}

func (uc *CategoryDeleteUseCase) Execute(ctx context.Context, input CategoryDeleteInput) error {
	err := uc.validate.Struct(input)
	if err != nil {
		return err
	}

	criteria := repository.DeleteCategoryCriteria{
		UserID:     input.UserID,
		CategoryID: input.CategoryID,
	}

	err = uc.categoryRepository.Delete(ctx, criteria)
	if err != nil {
		return err
	}

	return nil
}
