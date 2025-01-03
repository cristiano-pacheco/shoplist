package usecase

import "go.uber.org/fx"

var Module = fx.Module(
	"usecase",
	fx.Provide(NewCategoryCreateUseCase),
	fx.Provide(NewCategoryDeleteUseCase),
	fx.Provide(NewCategoryFindUseCase),
	fx.Provide(NewCategoryUpdateUseCase),
)
