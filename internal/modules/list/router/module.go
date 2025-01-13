package router

import "go.uber.org/fx"

var Module = fx.Module(
	"list/router",
	fx.Provide(
		NewRouter,
	),
	fx.Invoke(
		RegisterCategoryRoutes,
	),
)
