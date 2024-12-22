package rabbitmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ interface {
	Connection() *amqp.Connection
	Close()
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	VHost    string
}

func New(cfg Config) RabbitMQ {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.VHost,
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal("failed to connect to RabbitMQ", "error", err)
	}

	return &rabbitMQ{conn}
}

// facade represents the internal implementation of the Facade interface
type rabbitMQ struct {
	connection *amqp.Connection
}

func (f *rabbitMQ) Connection() *amqp.Connection {
	return f.connection
}

func (f *rabbitMQ) Close() {
	if err := f.connection.Close(); err != nil {
		fmt.Println("failed to close connection: %w", err)
	}
}
