package consumer

import (
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/queue"
	"github.com/cristiano-pacheco/shoplist/pkg/rabbitmq"
)

type SendUserConfirmationEmailConsumer interface {
	Start()
}

type sendUserConfirmationEmailConsumer struct {
	rabbitMQ rabbitmq.RabbitMQ
}

func NewSendUserConfirmationEmailConsumer(rabbitMQ rabbitmq.RabbitMQ) SendUserConfirmationEmailConsumer {
	return &sendUserConfirmationEmailConsumer{rabbitMQ}
}

func (c *sendUserConfirmationEmailConsumer) Start() {
	channel, err := c.rabbitMQ.Connection().Channel()
	if err != nil {
		return
	}
	defer channel.Close()
	err = c.rabbitMQ.DeclareDirectQueue(queue.SendUserConfirmationEmailQueue)
	if err != nil {
		return
	}
}
