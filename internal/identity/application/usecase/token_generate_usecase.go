package usecase

import (
	"context"
	"errors"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/service"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type TokenGenerateUseCase interface {
	Execute(ctx context.Context, input TokenGenerateInput) (TokenGenerateOutput, error)
}

type TokenGenerateInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type TokenGenerateOutput struct {
	Token string
}

type tokenGenerateUseCase struct {
	validator    validator.Validate
	userRepo     repository.UserRepository
	hashService  service.HashService
	tokenService service.TokenService
}

func NewTokenGenerateUseCase(
	validator validator.Validate,
	userRepo repository.UserRepository,
	hashService service.HashService,
	tokenService service.TokenService,
) TokenGenerateUseCase {
	return &tokenGenerateUseCase{
		validator,
		userRepo,
		hashService,
		tokenService,
	}
}

func (uc *tokenGenerateUseCase) Execute(ctx context.Context, input TokenGenerateInput) (TokenGenerateOutput, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "TokenGenerateUseCase.Execute")
	defer span.End()

	output := TokenGenerateOutput{}

	err := uc.validator.Struct(input)
	if err != nil {
		return output, err
	}

	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return output, errs.ErrInvalidCredentials
		}
		return output, err
	}

	if !user.IsActivated() {
		return output, errs.ErrUserIsNotActivated
	}

	hash := []byte(user.PasswordHash())
	pass := []byte(input.Password)
	err = uc.hashService.CompareHashAndPassword(hash, pass)
	if err != nil {
		return output, errs.ErrInvalidCredentials
	}

	token, err := uc.tokenService.Generate(ctx, user)
	if err != nil {
		return output, err
	}

	output.Token = token
	return output, nil
}
