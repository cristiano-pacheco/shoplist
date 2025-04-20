package handler

import "go.uber.org/fx"

var Module = fx.Module(
	"list/handler",
	fx.Provide(
		NewCategoryHandler,
		NewListHandler,
	),
)
