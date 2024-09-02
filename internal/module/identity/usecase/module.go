package usecase

import "go.uber.org/fx"

var Module = fx.Module(
	"billing/usecase",
	fx.Provide(NewCreateBillingUseCase),
)
