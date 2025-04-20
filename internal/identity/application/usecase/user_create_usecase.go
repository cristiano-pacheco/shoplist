package usecase

import (
	"context"
	"errors"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/errs"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/service"
	shared_errs "github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type UserCreateUseCase interface {
	Execute(ctx context.Context, input UserCreateInput) (UserCreateOutput, error)
}

type UserCreateInput struct {
	Name     string `validate:"required,min=3,max=255"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type UserCreateOutput struct {
	Name   string
	Email  string
	UserID uint64
}

type userCreateUseCase struct {
	sendEmailConfirmationService service.SendEmailConfirmationService
	hashService                  service.HashService
	userRepo                     repository.UserRepository
	validate                     validator.Validate
	logger                       logger.Logger
}

func NewUserCreateUseCase(
	sendEmailConfirmationService service.SendEmailConfirmationService,
	hashService service.HashService,
	userRepo repository.UserRepository,
	validate validator.Validate,
	logger logger.Logger,
) UserCreateUseCase {
	return &userCreateUseCase{
		sendEmailConfirmationService,
		hashService,
		userRepo,
		validate,
		logger,
	}
}

func (uc *userCreateUseCase) Execute(ctx context.Context, input UserCreateInput) (UserCreateOutput, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserCreateUseCase.Execute")
	defer span.End()

	output := UserCreateOutput{}

	err := uc.validate.Struct(input)
	if err != nil {
		return output, err
	}

	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, shared_errs.ErrNotFound) {
		uc.logger.Error("error finding user by email", "error", err)
		return output, err
	}

	if user.ID() != 0 {
		return output, errs.ErrEmailAlreadyInUse
	}

	ph, err := uc.hashService.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		message := "error generating password hash"
		uc.logger.Error(message, "error", err)
		return output, err
	}

	userModel, err := model.CreateUserModel(input.Name, input.Email, string(ph))
	if err != nil {
		message := "error creating user model"
		uc.logger.Error(message, "error", err)
		return output, err
	}

	newUserModel, err := uc.userRepo.Create(ctx, userModel)
	if err != nil {
		message := "error creating user"
		uc.logger.Error(message, "error", err)
		return output, err
	}

	err = uc.sendEmailConfirmationService.Execute(ctx, uint64(newUserModel.ID()))
	if err != nil {
		message := "error sending account confirmation email"
		uc.logger.Error(message, "error", err)
		return output, err
	}

	output = UserCreateOutput{
		UserID: uint64(newUserModel.ID()),
		Name:   newUserModel.Name(),
		Email:  newUserModel.Email(),
	}

	return output, nil
}
