package category_delete

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type CategoryDeleteUseCaseI interface {
	Execute(ctx context.Context, input Input) error
}

type CategoryDeleteUseCase struct {
	categoryRepository repository.CategoryRepositoryI
	validate           validator.ValidateI
}

func New(categoryRepository repository.CategoryRepositoryI, validate validator.ValidateI) CategoryDeleteUseCaseI {
	return &CategoryDeleteUseCase{categoryRepository, validate}
}

func (uc *CategoryDeleteUseCase) Execute(ctx context.Context, input Input) error {
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
