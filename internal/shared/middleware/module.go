package middleware

import (
	"github.com/cristiano-pacheco/go-modulith/internal/shared/middleware/auth_middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared/middleware",
	auth_middleware.Module,
)
