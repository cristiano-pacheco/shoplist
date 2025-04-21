package mapper

import (
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/entity"
)

type AccountConfirmationMapper interface {
	ToModel(entity entity.AccountConfirmationEntity) (model.AccountConfirmationModel, error)
	ToEntity(model model.AccountConfirmationModel) entity.AccountConfirmationEntity
}

type accountConfirmationMapper struct {
}

func NewAccountConfirmationMapper() AccountConfirmationMapper {
	return &accountConfirmationMapper{}
}

func (u *accountConfirmationMapper) ToModel(entity entity.AccountConfirmationEntity) (model.AccountConfirmationModel, error) {
	accountConfirmationModel, err := model.RestoreAccountConfirmationModel(
		entity.ID,
		entity.UserID,
		entity.Token,
		entity.ExpiresAt,
		entity.CreatedAt,
	)
	if err != nil {
		return model.AccountConfirmationModel{}, err
	}
	return accountConfirmationModel, nil
}

func (u *accountConfirmationMapper) ToEntity(model model.AccountConfirmationModel) entity.AccountConfirmationEntity {
	return entity.AccountConfirmationEntity{
		ID:        model.ID(),
		UserID:    model.UserID(),
		Token:     model.Token(),
		ExpiresAt: model.ExpiresAt(),
		CreatedAt: model.CreatedAt(),
	}
}
