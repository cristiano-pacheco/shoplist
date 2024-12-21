package consumer

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service"
)

type UserConfirmationEmailConsumer interface {
	Execute(ctx context.Context, message dto.SendConfirmationEmailMessage) error
}

type userConfirmationEmailConsumer struct {
	userRepo                 repository.UserRepository
	emailConfirmationService service.EmailConfirmationService
}

func NewUserConfirmationEmailConsumer(
	userRepo repository.UserRepository,
	emailConfirmationService service.EmailConfirmationService,
) UserConfirmationEmailConsumer {
	return &userConfirmationEmailConsumer{userRepo, emailConfirmationService}
}

func (c *userConfirmationEmailConsumer) Execute(
	ctx context.Context,
	message dto.SendConfirmationEmailMessage,
) error {
	err := c.emailConfirmationService.Send(ctx, message.UserID)
	if err != nil {
		return err
	}

	return nil
}
