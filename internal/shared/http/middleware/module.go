package middleware

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/middleware/auth_middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared/middleware",
	auth_middleware.Module,
	fx.Provide(NewErrorHandlerMiddleware),
)
