package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/list/repository"
)

type ListFindUsecase struct {
	listRepository repository.ListRepository
}

func NewListFindUsecase(listRepository repository.ListRepository) *ListFindUsecase {
	return &ListFindUsecase{listRepository}
}

func (uc *ListFindUsecase) ExecuteByIDAndUserID(ctx context.Context, id uint64, userID uint64) (model.ListModel, error) {
	list, err := uc.listRepository.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return model.ListModel{}, err
	}

	return list, nil
}

func (uc *ListFindUsecase) ExecuteByUserID(ctx context.Context, userID uint64) ([]model.ListModel, error) {
	list, err := uc.listRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return list, nil
}
