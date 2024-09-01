package repository

import "go.uber.org/fx"

var Module = fx.Module(
	"billing/repository",
	fx.Provide(NewBillingRepository),
	fx.Provide(NewCategoryRepository),
)
