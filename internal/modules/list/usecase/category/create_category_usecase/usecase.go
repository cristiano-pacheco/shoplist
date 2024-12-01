package create_category_usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type UseCaseI interface {
	Execute(ctx context.Context, input Input) (Output, error)
}

type useCase struct {
	categoryRepository repository.CategoryRepositoryI
	validate           validator.ValidateI
}

func New(
	categoryRepository repository.CategoryRepositoryI,
	validate validator.ValidateI,
) UseCaseI {
	return &useCase{categoryRepository, validate}
}

func (uc *useCase) Execute(ctx context.Context, input Input) (Output, error) {
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
