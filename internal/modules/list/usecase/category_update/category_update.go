package category_update

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type CategoryUpdateUseCaseI interface {
	Execute(ctx context.Context, input Input) error
}

type CategoryUpdateUseCase struct {
	categoryRepository repository.CategoryRepositoryI
	validate           validator.ValidateI
}

func New(categoryRepository repository.CategoryRepositoryI, validate validator.ValidateI) CategoryUpdateUseCaseI {
	return &CategoryUpdateUseCase{categoryRepository, validate}
}

func (uc *CategoryUpdateUseCase) Execute(ctx context.Context, input Input) error {
	err := uc.validate.Struct(input)
	if err != nil {
		return err
	}

	categoryModel := model.CategoryModel{
		ID:     input.CategoryID,
		UserID: input.UserID,
		Name:   input.Name,
	}

	err = uc.categoryRepository.Update(ctx, categoryModel)
	if err != nil {
		return err
	}

	return nil
}
