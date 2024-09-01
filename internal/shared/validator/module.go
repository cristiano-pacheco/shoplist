package validator

import "go.uber.org/fx"

var Module = fx.Module("shared/validator", fx.Provide(New))
