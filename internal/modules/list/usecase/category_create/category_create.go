package category_create

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type CategoryCreateUseCaseI interface {
	Execute(ctx context.Context, input Input) (Output, error)
}

type CategoryCreateUseCase struct {
	categoryRepository repository.CategoryRepositoryI
	validate           validator.ValidateI
}

func New(
	categoryRepository repository.CategoryRepositoryI,
	validate validator.ValidateI,
) CategoryCreateUseCaseI {
	return &CategoryCreateUseCase{categoryRepository, validate}
}

func (uc *CategoryCreateUseCase) Execute(ctx context.Context, input Input) (Output, error) {
	err := uc.validate.Struct(input)
	if err != nil {
		return Output{}, err
	}

	categoryModel := model.CategoryModel{
		UserID: input.UserID,
		Name:   input.Name,
	}

	dbCategoryModel, err := uc.categoryRepository.Create(ctx, categoryModel)
	if err != nil {
		return Output{}, err
	}

	return Output{CategoryModel: *dbCategoryModel}, nil
}
