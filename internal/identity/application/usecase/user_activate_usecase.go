package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/errs"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/logger"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/otel"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/validator"
)

type UserActivateUseCase interface {
	Execute(ctx context.Context, input UserActivateInput) error
}

type UserActivateInput struct {
	Token string `validate:"required"`
}

type userActivateUseCase struct {
	userRepository repository.UserRepository
	validate       validator.Validate
	logger         logger.Logger
}

func NewUserActivateUseCase(
	userRepository repository.UserRepository,
	validate validator.Validate,
	logger logger.Logger,
) UserActivateUseCase {
	return &userActivateUseCase{userRepository, validate, logger}
}

func (uc *userActivateUseCase) Execute(ctx context.Context, input UserActivateInput) error {
	ctx, span := otel.Trace().StartSpan(ctx, "UserActivateUseCase.Execute")
	defer span.End()

	err := uc.validate.Struct(input)
	if err != nil {
		return err
	}

	user, err := uc.userRepository.FindByConfirmationToken(ctx, input.Token)
	if err != nil {
		return err
	}

	if !user.IsConfirmationTokenValid(input.Token) {
		return errs.ErrInvalidAccountConfirmationToken
	}

	user.ConfirmAccount()
	err = uc.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
