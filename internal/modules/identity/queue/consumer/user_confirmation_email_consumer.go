package consumer

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service"
	"github.com/cristiano-pacheco/shoplist/pkg/rabbitmq"
)

type UserConfirmationEmailConsumer interface {
	Start() error
	WaitForCompletion()
}

type userConfirmationEmailConsumer struct {
	userRepo                 repository.UserRepository
	emailConfirmationService service.EmailConfirmationService
	rabbitMQ                 rabbitmq.Facade
	wg                       sync.WaitGroup
	processingComplete       chan struct{}
}

func NewUserConfirmationEmailConsumer(
	userRepo repository.UserRepository,
	emailConfirmationService service.EmailConfirmationService,
	rabbitMQ rabbitmq.Facade,
) UserConfirmationEmailConsumer {
	return &userConfirmationEmailConsumer{
		userRepo:                 userRepo,
		emailConfirmationService: emailConfirmationService,
		rabbitMQ:                 rabbitMQ,
		processingComplete:       make(chan struct{}),
	}
}

func (c *userConfirmationEmailConsumer) Start() error {
	queueName := "user-confirmation-email"
	log.Printf("Starting consumer for queue: %s\n", queueName)

	messages, err := c.rabbitMQ.ConsumeWithChannel(queueName)
	if err != nil {
		return err
	}

	go func() {
		for msg := range messages {
			c.wg.Add(1)
			go func(delivery rabbitmq.Delivery) {
				defer c.wg.Done()

				var message dto.SendConfirmationEmailMessage
				if err := json.Unmarshal(delivery.Body(), &message); err != nil {
					log.Printf("Error unmarshaling message: %v", err)
					delivery.Nack(false, false)
					return
				}

				if err := c.execute(context.Background(), message); err != nil {
					log.Printf("Error processing message: %v", err)
					delivery.Nack(false, true)
					return
				}

				delivery.Ack(false)
			}(msg)
		}

		c.wg.Wait()
		close(c.processingComplete)
	}()

	return nil
}

func (c *userConfirmationEmailConsumer) WaitForCompletion() {
	<-c.processingComplete
}

func (c *userConfirmationEmailConsumer) execute(
	ctx context.Context,
	message dto.SendConfirmationEmailMessage,
) error {
	err := c.emailConfirmationService.Send(ctx, message.UserID)
	if err != nil {
		return err
	}

	return nil
}
