package model

import "github.com/cristiano-pacheco/shoplist/internal/shared/model"

type UserModel struct {
	model.Base
	Name         string `gorm:"type:varchar;not null; column:name"`
	Email        string `gorm:"type:varchar;not null; column:email"`
	PasswordHash string `gorm:"type:varchar;not null; column:password_hash"`
	IsActivated  bool   `gorm:"type:boolean;not null;default:false;column:is_activated"`
}

func (UserModel) TableName() string {
	return "users"
}
