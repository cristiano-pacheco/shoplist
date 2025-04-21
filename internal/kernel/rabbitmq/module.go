package rabbitmq

import "go.uber.org/fx"

var Module = fx.Module("kernel/rabbitmq", fx.Provide(New))
