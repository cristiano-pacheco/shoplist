package mapper

import (
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mapper/errormapper"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared/mapper",
	fx.Provide(errormapper.New),
)
