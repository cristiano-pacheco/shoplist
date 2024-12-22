package producer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/queue"
	"github.com/cristiano-pacheco/shoplist/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

type UserConfirmationEmailProducer interface {
	Execute(ctx context.Context, message dto.SendConfirmationEmailMessage) error
}

type userConfirmationEmailProducer struct {
	rabbitMQ rabbitmq.RabbitMQ
}

func NewUserConfirmationEmailProducer(rabbitMQ rabbitmq.RabbitMQ) UserConfirmationEmailProducer {
	return &userConfirmationEmailProducer{rabbitMQ}
}

func (p *userConfirmationEmailProducer) Execute(
	ctx context.Context,
	message dto.SendConfirmationEmailMessage,
) error {
	m, err := json.Marshal(message)
	if err != nil {
		return err
	}

	channel, err := p.rabbitMQ.Connection().Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = p.declareQueue(channel)
	if err != nil {
		return err
	}

	err = channel.Publish(
		queue.SendUserConfirmationEmailQueue, // exchange
		"",                                   // routing key
		false,                                // mandatory
		false,                                // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Timestamp:   time.Now(),
			Body:        m,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *userConfirmationEmailProducer) declareQueue(channel *amqp091.Channel) error {
	err := channel.ExchangeDeclare(
		queue.SendUserConfirmationEmailQueue,
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	_, err = channel.QueueDeclare(
		queue.SendUserConfirmationEmailQueue,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = channel.QueueBind(
		queue.SendUserConfirmationEmailQueue,
		"",
		queue.SendUserConfirmationEmailQueue,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
