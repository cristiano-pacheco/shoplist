package rabbitmq

import "go.uber.org/fx"

var Module = fx.Module("shared/rabbitmq", fx.Provide(New))
