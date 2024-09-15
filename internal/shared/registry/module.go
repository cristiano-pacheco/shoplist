package registry

import (
	"github.com/cristiano-pacheco/go-modulith/internal/shared/registry/privatekey_registry"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared/registry",
	fx.Provide(privatekey_registry.New),
)
