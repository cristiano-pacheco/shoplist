package errormapper

import "go.uber.org/fx"

var Module = fx.Module("errormapper", fx.Provide(New))
