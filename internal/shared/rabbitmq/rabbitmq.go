package rabbitmq

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/pkg/rabbitmq"
	"go.uber.org/fx"
)

type Facade rabbitmq.Facade

func New(lc fx.Lifecycle, cfg config.Config) Facade {
	rabbitMQConfig := rabbitmq.Config{
		Host:     cfg.RabbitMQ.Host,
		Port:     cfg.RabbitMQ.Port,
		Username: cfg.RabbitMQ.Username,
		Password: cfg.RabbitMQ.Password,
		VHost:    cfg.RabbitMQ.VHost,
	}

	facade := rabbitmq.New(rabbitMQConfig)

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			facade.Close()
			return nil
		},
	})

	return Facade(facade)
}
