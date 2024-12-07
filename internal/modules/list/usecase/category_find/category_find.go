package category_find

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
)

type CategoryFindUseCase struct {
	categoryRepo repository.CategoryRepository
}

func New(categoryRepo repository.CategoryRepository) *CategoryFindUseCase {
	return &CategoryFindUseCase{categoryRepo}
}

func (uc *CategoryFindUseCase) Execute(ctx context.Context, input Input) (Output, error) {
	criteria := repository.FindCategoriesCriteria{
		UserID:     input.UserID,
		CategoryID: input.CategoryID,
		Name:       input.CategoryName,
	}

	categories, err := uc.categoryRepo.Find(ctx, criteria)
	if err != nil {
		return Output{}, err
	}

	return Output{Categories: categories}, nil
}
