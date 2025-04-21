package mapper

import (
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/entity"
)

type UserMapper interface {
	ToModel(entity entity.UserEntity) (model.UserModel, error)
	ToEntity(model model.UserModel) entity.UserEntity
}

type userMapper struct {
}

func NewUserMapper() UserMapper {
	return &userMapper{}
}

func (u *userMapper) ToModel(entity entity.UserEntity) (model.UserModel, error) {
	userModel, err := model.RestoreUserModel(
		entity.ID,
		entity.Name,
		entity.Email,
		entity.PasswordHash,
		entity.IsActivated,
		entity.RpToken,
		entity.CreatedAt,
		entity.UpdatedAt,
	)
	if err != nil {
		return model.UserModel{}, err
	}
	return userModel, nil
}

func (u *userMapper) ToEntity(model model.UserModel) entity.UserEntity {
	return entity.UserEntity{
		ID:           model.ID(),
		Name:         model.Name(),
		Email:        model.Email(),
		PasswordHash: model.PasswordHash(),
		IsActivated:  model.IsActivated(),
		RpToken:      model.RpToken(),
		CreatedAt:    model.CreatedAt(),
		UpdatedAt:    model.UpdatedAt(),
	}
}
