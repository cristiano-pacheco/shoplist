package rabbitmq

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/kernel/config"
	"github.com/cristiano-pacheco/shoplist/pkg/rabbitmq"
	"go.uber.org/fx"
)

func New(lc fx.Lifecycle, cfg config.Config) rabbitmq.RabbitMQ {
	rabbitMQConfig := rabbitmq.Config{
		Host:     cfg.RabbitMQ.Host,
		Port:     cfg.RabbitMQ.Port,
		Username: cfg.RabbitMQ.Username,
		Password: cfg.RabbitMQ.Password,
		VHost:    cfg.RabbitMQ.VHost,
	}

	rabbitMQ := rabbitmq.New(rabbitMQConfig)

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			rabbitMQ.Close()
			return nil
		},
	})

	return rabbitMQ
}
