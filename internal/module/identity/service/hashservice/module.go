package hashservice

import "go.uber.org/fx"

var Module = fx.Module("hashservice", fx.Provide(New))
