package producer

import (
	"context"
	"encoding/json"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/dto"
	"github.com/cristiano-pacheco/shoplist/pkg/rabbitmq"
)

type UserConfirmationEmailProducer interface {
	Execute(ctx context.Context, message dto.SendConfirmationEmailMessage) error
}

type userConfirmationEmailProducer struct {
	rabbitMQ rabbitmq.Facade
}

func NewUserConfirmationEmailProducer(rabbitMQ rabbitmq.Facade) UserConfirmationEmailProducer {
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

	return p.rabbitMQ.Publish(ctx, "user-confirmation-email", m)
}
