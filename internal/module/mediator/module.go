package mediator

import "go.uber.org/fx"

var Module = fx.Module("mediator", fx.Provide(New))
