package model

import "github.com/cristiano-pacheco/go-modulith/internal/shared/model"

type CategoryModel struct {
	model.Base
	Name string `gorm:"type:varchar;not null; column:name"`
}

func (CategoryModel) TableName() string {
	return "category"
}
