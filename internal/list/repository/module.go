package repository

import "go.uber.org/fx"

var Module = fx.Module(
	"list/repository",
	fx.Provide(
		NewCategoryRepository,
	),
)
