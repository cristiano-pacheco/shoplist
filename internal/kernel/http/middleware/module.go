package middleware

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"kernel/middleware",
	fx.Provide(NewErrorHandlerMiddleware),
)
