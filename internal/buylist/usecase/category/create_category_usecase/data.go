package create_category_usecase

import "github.com/cristiano-pacheco/go-modulith/internal/buylist/model"

type Input struct {
	UserID uint64 `validate:"required"`
	Name   string `validate:"required,min=1,max=255"`
}

type Output struct {
	CategoryModel model.CategoryModel
}
