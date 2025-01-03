package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
)

type CategoryFindUseCase struct {
	categoryRepo repository.CategoryRepository
}

type CategoryFindInput struct {
	UserID       uint64
	CategoryID   *uint64
	CategoryName *string
}

type CategoryFindOutput struct {
	CategoryList []*model.CategoryModel
}

func NewCategoryFindUseCase(categoryRepo repository.CategoryRepository) *CategoryFindUseCase {
	return &CategoryFindUseCase{categoryRepo}
}

func (uc *CategoryFindUseCase) Execute(ctx context.Context, input CategoryFindInput) (CategoryFindOutput, error) {
	criteria := repository.FindCategoriesCriteria{
		UserID:     input.UserID,
		CategoryID: input.CategoryID,
		Name:       input.CategoryName,
	}

	categories, err := uc.categoryRepo.Find(ctx, criteria)
	if err != nil {
		return CategoryFindOutput{}, err
	}

	return CategoryFindOutput{CategoryList: categories}, nil
}
