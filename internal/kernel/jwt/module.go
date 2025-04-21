package jwt

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared/jwt",
	fx.Provide(NewParser),
)
