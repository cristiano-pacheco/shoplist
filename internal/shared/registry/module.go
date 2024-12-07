package registry

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared/registry",
	fx.Provide(
		NewPrivateKeyRegistry,
	),
)
