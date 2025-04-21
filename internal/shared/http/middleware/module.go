package middleware

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared/middleware",
	fx.Provide(NewErrorHandlerMiddleware),
)
