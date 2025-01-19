package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/repository"
)

type ListFindUseCase struct {
	listRepository repository.ListRepository
}

type ListFindInput struct {
	UserID   uint64
	ListID   *uint64
	ListName *string
}

type ListFindOutput struct {
	Lists []model.ListModel
}

func NewListFindUseCase(listRepository repository.ListRepository) *ListFindUseCase {
	return &ListFindUseCase{listRepository}
}

func (uc *ListFindUseCase) Execute(ctx context.Context, input ListFindInput) (ListFindOutput, error) {
	criteria := repository.FindListsCriteria{
		UserID: input.UserID,
		ListID: input.ListID,
		Name:   input.ListName,
	}

	lists, err := uc.listRepository.Find(ctx, criteria)
	if err != nil {
		return ListFindOutput{}, err
	}

	return ListFindOutput{Lists: lists}, nil
}
