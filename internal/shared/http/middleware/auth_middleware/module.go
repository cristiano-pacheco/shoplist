package auth_middleware

import "go.uber.org/fx"

var Module = fx.Module(
	"shared/middleware/auth_middleware",
	fx.Provide(New),
	fx.Provide(newIsUserEnabledQuery),
)
