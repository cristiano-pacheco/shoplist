package rabbitmq

import (
	"log"

	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/pkg/rabbitmq"
)

type Facade rabbitmq.Facade

func New(cfg config.Config) Facade {
	rabbitMQConfig := rabbitmq.Config{
		Host:     cfg.RabbitMQ.Host,
		Port:     cfg.RabbitMQ.Port,
		Username: cfg.RabbitMQ.Username,
		Password: cfg.RabbitMQ.Password,
		VHost:    cfg.RabbitMQ.VHost,
	}
	facade, err := rabbitmq.New(rabbitMQConfig)
	if err != nil {
		log.Fatalf("Error creating RabbitMQ facade: %v", err)
	}
	return Facade(facade)
}
