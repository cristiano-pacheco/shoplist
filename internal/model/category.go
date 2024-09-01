package model

type CategoryModel struct {
	Base
	Name string `gorm:"type:varchar;not null; column:name"`
}

func (CategoryModel) TableName() string {
	return "category"
}
