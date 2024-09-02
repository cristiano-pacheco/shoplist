package handler

import "go.uber.org/fx"

var Module = fx.Module(
	"billing/handler",
	fx.Provide(NewBillingHandler),
)
