package mapper

import (
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/entity"
)

type LoginTokenMapper interface {
	ToModel(entity entity.LoginTokenEntity) (model.LoginTokenModel, error)
	ToEntity(model model.LoginTokenModel) entity.LoginTokenEntity
}

type loginTokenMapper struct {
}

func NewLoginTokenMapper() LoginTokenMapper {
	return &loginTokenMapper{}
}

func (m *loginTokenMapper) ToModel(entity entity.LoginTokenEntity) (model.LoginTokenModel, error) {
	loginTokenModel, err := model.RestoreLoginTokenModel(
		entity.ID,
		entity.UserID,
		entity.Token,
		entity.ExpiresAt,
		entity.CreatedAt,
		entity.UpdatedAt,
	)
	if err != nil {
		return model.LoginTokenModel{}, err
	}
	return loginTokenModel, nil
}

func (m *loginTokenMapper) ToEntity(model model.LoginTokenModel) entity.LoginTokenEntity {
	return entity.LoginTokenEntity{
		ID:        model.ID(),
		UserID:    model.UserID(),
		Token:     model.Token(),
		ExpiresAt: model.ExpiresAt(),
		CreatedAt: model.CreatedAt(),
		UpdatedAt: model.UpdatedAt(),
	}
}
