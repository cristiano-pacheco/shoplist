package find_category_usecase

import "github.com/cristiano-pacheco/shoplist/internal/modules/list/model"

type Input struct {
	UserID uint64
}

type Output struct {
	Categories []model.CategoryModel
}
