package usecase

import (
	"context"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	"github.com/cristiano-pacheco/shoplist/internal/shared/sdk/empty"
)

type UserActivateUseCase interface {
	Execute(ctx context.Context, input UserActivateUseCaseInput) error
}

type UserActivateUseCaseInput struct {
	UserID uint64 `validate:"required,number"`
	Token  string `validate:"required"`
}

type userActivateUseCase struct {
	userRepo                repository.UserRepository
	accountConfirmationRepo repository.AccountConfirmationRepository
	logger                  logger.Logger
}

func NewUserActivateUseCase(
	userRepo repository.UserRepository,
	accountConfirmationRepo repository.AccountConfirmationRepository,
	logger logger.Logger,
) UserActivateUseCase {
	return &userActivateUseCase{userRepo, accountConfirmationRepo, logger}
}

func (uc *userActivateUseCase) Execute(ctx context.Context, input UserActivateUseCaseInput) error {
	ctx, span := otel.Trace().StartSpan(ctx, "UserActivateUseCase.Execute")
	defer span.End()

	err := uc.validateInput(input)
	if err != nil {
		uc.logger.InfoContext(ctx, "invalid input", "error", err)
		return err
	}

	accountConfirmationModel, err := uc.accountConfirmationRepo.FindByUserID(ctx, input.UserID)
	if err != nil {
		uc.logger.ErrorContext(
			ctx,
			"error finding account confirmation with user_id",
			"error", err,
			"user_id", input.UserID,
		)
		return err
	}

	if accountConfirmationModel.Token() != input.Token {
		return errs.ErrInvalidAccountConfirmationToken
	}

	userModel, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		return err
	}

	userModel.Activate()
	err = uc.userRepo.Update(ctx, userModel)
	if err != nil {
		uc.logger.ErrorContext(ctx, "error updating user", "error", err)
		return err
	}

	err = uc.accountConfirmationRepo.DeleteById(ctx, accountConfirmationModel.ID())
	if err != nil {
		uc.logger.ErrorContext(ctx, "error deleting account confirmation", "error", err)
		return err
	}

	return nil
}

func (uc *userActivateUseCase) validateInput(input UserActivateUseCaseInput) error {
	if empty.IsEmpty(input.UserID) {
		return errs.NewBadRequestError("user_id is required")
	}

	if empty.IsEmpty(input.Token) {
		return errs.NewBadRequestError("token is required")
	}

	return nil
}
