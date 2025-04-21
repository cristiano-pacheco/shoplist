package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type CategoryUpdateUseCase struct {
	categoryRepository repository.CategoryRepository
	validate           validator.Validate
}

type CategoryUpdateInput struct {
	UserID     uint64 `validate:"required"`
	CategoryID uint64 `validate:"required"`
	Name       string `validate:"required,min=1,max=255"`
}

func NewCategoryUpdateUseCase(
	categoryRepository repository.CategoryRepository,
	validate validator.Validate,
) *CategoryUpdateUseCase {
	return &CategoryUpdateUseCase{categoryRepository, validate}
}

func (uc *CategoryUpdateUseCase) Execute(ctx context.Context, input CategoryUpdateInput) error {
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
