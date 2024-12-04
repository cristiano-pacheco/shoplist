package category_find

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
)

type CategoryFindUseCaseI interface {
	Execute(ctx context.Context, input Input) (Output, error)
}

type CategoryFindUseCase struct {
	categoryRepo repository.CategoryRepositoryI
}

func New(categoryRepo repository.CategoryRepositoryI) CategoryFindUseCaseI {
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
