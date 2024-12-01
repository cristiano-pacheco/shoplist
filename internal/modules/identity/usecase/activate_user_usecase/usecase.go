package activate_user_usecase

import (
	"context"
	"time"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
)

type UseCase struct {
	userRepo                repository.UserRepositoryI
	accountConfirmationRepo repository.AccountConfirmationRepositoryI
	logger                  logger.LoggerI
}

func New(
	userRepo repository.UserRepositoryI,
	accountConfirmationRepo repository.AccountConfirmationRepositoryI,
	logger logger.LoggerI,
) *UseCase {
	return &UseCase{userRepo, accountConfirmationRepo, logger}
}

func (u *UseCase) Execute(ctx context.Context, input Input) error {
	err := validateInput(input)
	if err != nil {
		u.logger.InfoContext(ctx, "[activate_user_usecase] invalid input", "error", err)
		return err
	}

	accountConfirmationModel, err := u.accountConfirmationRepo.FindByUserID(ctx, input.UserID)
	if err != nil {
		u.logger.ErrorContext(
			ctx,
			"[activate_user_usecase] error finding account confirmation with user_id",
			"error", err,
			"user_id", input.UserID,
		)
		return err
	}

	if accountConfirmationModel.Token != input.Token {
		return errs.ErrInvalidAccountConfirmationToken
	}

	userModel, err := u.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		return err
	}

	userModel.IsActivated = true
	userModel.UpdatedAt = time.Now().UTC()
	err = u.userRepo.Update(ctx, *userModel)
	if err != nil {
		u.logger.ErrorContext(ctx, "[activate_user_usecase] error updating user", "error", err)
		return err
	}

	err = u.accountConfirmationRepo.Delete(ctx, accountConfirmationModel)
	if err != nil {
		u.logger.ErrorContext(ctx, "[activate_user_usecase] error deleting account confirmation", "error", err)
		return err
	}

	return nil
}
