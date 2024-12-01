package create_category_usecase

import "github.com/cristiano-pacheco/shoplist/internal/modules/list/model"

type Input struct {
	UserID uint64 `validate:"required"`
	Name   string `validate:"required,min=1,max=255"`
}

type Output struct {
	CategoryModel model.CategoryModel
}
