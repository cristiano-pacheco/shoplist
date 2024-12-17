package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Facade represents the interface for RabbitMQ operations
type Facade interface {
	// PublishMessage publishes a message to a specified exchange and routing key
	PublishMessage(ctx context.Context, exchange, routingKey string, message []byte) error
	// ConsumeMessages starts consuming messages from a specified queue
	ConsumeMessages(queueName string, handler func([]byte) error) error
	// Close closes the RabbitMQ connection and channel
	Close() error
}

// Config holds the configuration for RabbitMQ connection
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	VHost    string
}

// New creates a new instance of RabbitMQ Facade
func New(cfg Config) (Facade, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.VHost,
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	return &facade{
		conn:    conn,
		channel: channel,
	}, nil
}

// facade represents the internal implementation of the Facade interface
type facade struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func (f *facade) PublishMessage(ctx context.Context, exchange, routingKey string, message []byte) error {
	err := f.channel.PublishWithContext(
		ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

func (f *facade) ConsumeMessages(queueName string, handler func([]byte) error) error {
	msgs, err := f.channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	go func() {
		for msg := range msgs {
			err := handler(msg.Body)
			if err != nil {
				// If handler returns an error, message will be nacked and requeued
				msg.Nack(false, true)
				continue
			}
			// If handler succeeds, message will be acked
			msg.Ack(false)
		}
	}()

	return nil
}

func (f *facade) Close() error {
	if err := f.channel.Close(); err != nil {
		return fmt.Errorf("failed to close channel: %w", err)
	}
	if err := f.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}
	return nil
}
