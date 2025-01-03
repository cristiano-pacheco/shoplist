package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type CategoryCreateUseCase struct {
	categoryRepository repository.CategoryRepository
	validate           validator.Validate
}

type CategoryCreateInput struct {
	UserID uint64 `validate:"required"`
	Name   string `validate:"required,min=1,max=255"`
}

type CategoryCreateOutput struct {
	CategoryModel model.CategoryModel
}

func NewCategoryCreateUseCase(
	categoryRepository repository.CategoryRepository,
	validate validator.Validate,
) *CategoryCreateUseCase {
	return &CategoryCreateUseCase{categoryRepository, validate}
}

func (uc *CategoryCreateUseCase) Execute(ctx context.Context, input CategoryCreateInput) (CategoryCreateOutput, error) {
	err := uc.validate.Struct(input)
	if err != nil {
		return CategoryCreateOutput{}, err
	}

	categoryModel := model.CategoryModel{
		UserID: input.UserID,
		Name:   input.Name,
	}

	dbCategoryModel, err := uc.categoryRepository.Create(ctx, categoryModel)
	if err != nil {
		return CategoryCreateOutput{}, err
	}

	return CategoryCreateOutput{CategoryModel: *dbCategoryModel}, nil
}
