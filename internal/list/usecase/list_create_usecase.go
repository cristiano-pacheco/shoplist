package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/list/model"
	"github.com/cristiano-pacheco/shoplist/internal/list/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type ListCreateUseCase struct {
	listRepository repository.ListRepository
	validate       validator.Validate
}

type ListCreateInput struct {
	UserID uint64 `validate:"required"`
	Name   string `validate:"required,min=1,max=255"`
}

type ListCreateOutput struct {
	ListModel model.ListModel
}

func NewListCreateUseCase(
	listRepository repository.ListRepository,
	validate validator.Validate,
) *ListCreateUseCase {
	return &ListCreateUseCase{listRepository, validate}
}

func (uc *ListCreateUseCase) Execute(ctx context.Context, input ListCreateInput) (ListCreateOutput, error) {
	err := uc.validate.Struct(input)
	if err != nil {
		return ListCreateOutput{}, err
	}

	model := model.ListModel{UserID: input.UserID, Name: input.Name}
	model, err = uc.listRepository.Create(ctx, model)
	if err != nil {
		return ListCreateOutput{}, err
	}

	return ListCreateOutput{ListModel: model}, nil
}
