package router

import "go.uber.org/fx"

var Module = fx.Module("billing/router",
	fx.Provide(NewRouter),
	fx.Invoke(RegisterBillingHandlers),
)
