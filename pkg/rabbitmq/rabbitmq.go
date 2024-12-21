package rabbitmq

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Facade interface {
	Publish(ctx context.Context, queueName string, message []byte) error
	Consume(queueName string, handler func([]byte) error) error
	Close()
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	VHost    string
}

func New(cfg Config) Facade {
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

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open channel", "error", err)
	}

	return &facade{
		conn:    conn,
		channel: channel,
	}
}

// facade represents the internal implementation of the Facade interface
type facade struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func (f *facade) Publish(ctx context.Context, queueName string, message []byte) error {
	err := f.declareQueue(queueName)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	err = f.channel.PublishWithContext(
		ctx,
		queueName, // exchange
		"",        // routing key
		false,     // mandatory
		false,     // immediate
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

func (f *facade) Consume(queueName string, handler func([]byte) error) error {
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

func (f *facade) Close() {
	if err := f.channel.Close(); err != nil {
		fmt.Println("failed to close channel: %w", err)
	}
	if err := f.conn.Close(); err != nil {
		fmt.Println("failed to close connection: %w", err)
	}
}

func (f *facade) declareQueue(queueName string) error {
	// Declare the queue
	queue, err := f.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Declare the exchange
	err = f.channel.ExchangeDeclare(
		queueName, // name (same as queue)
		"direct",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Bind the queue to the exchange
	err = f.channel.QueueBind(
		queue.Name, // queue name
		"",         // routing key (empty for direct exchange)
		queueName,  // exchange
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue to exchange: %w", err)
	}

	return nil
}
