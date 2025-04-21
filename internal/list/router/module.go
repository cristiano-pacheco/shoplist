package router

import "go.uber.org/fx"

var Module = fx.Module(
	"list/router",
	fx.Provide(
		NewV1FiberRouter,
	),
	fx.Invoke(
		RegisterCategoryRoutes,
	),
)
