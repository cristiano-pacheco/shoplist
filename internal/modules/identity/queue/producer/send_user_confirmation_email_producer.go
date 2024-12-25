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

type SendUserConfirmationEmailProducer interface {
	Execute(ctx context.Context, message dto.SendConfirmationEmailMessage) error
}

type sendUserConfirmationEmailProducer struct {
	rabbitMQ rabbitmq.RabbitMQ
}

func NewSendUserConfirmationEmailProducer(rabbitMQ rabbitmq.RabbitMQ) SendUserConfirmationEmailProducer {
	return &sendUserConfirmationEmailProducer{rabbitMQ}
}

func (p *sendUserConfirmationEmailProducer) Execute(
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

	err = p.rabbitMQ.DeclareDirectQueue(queue.SendUserConfirmationEmailQueue)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(
		ctx,
		queue.SendUserConfirmationEmailQueue, // exchange
		"",                                   // routing key
		false,                                // mandatory
		false,                                // immediate
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "application/json",
			Timestamp:    time.Now(),
			Body:         m,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
