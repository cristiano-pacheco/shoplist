package category_find

import "github.com/cristiano-pacheco/shoplist/internal/modules/list/model"

type Input struct {
	UserID       uint64
	CategoryID   *uint64
	CategoryName *string
}

type Output struct {
	Categories []*model.CategoryModel
}
