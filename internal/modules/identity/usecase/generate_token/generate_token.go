package generate_token

import (
	"context"
	"errors"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/telemetry"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type GenerateTokenUseCase struct {
	validator            validator.ValidateI
	userRepo             repository.UserRepositoryI
	hashService          service.HashServiceI
	generateTokenService service.GenerateTokenServiceI
}

func New(
	validator validator.ValidateI,
	userRepo repository.UserRepositoryI,
	hashService service.HashServiceI,
	generateTokenService service.GenerateTokenServiceI,
) *GenerateTokenUseCase {
	return &GenerateTokenUseCase{
		validator,
		userRepo,
		hashService,
		generateTokenService,
	}
}

func (uc *GenerateTokenUseCase) Execute(ctx context.Context, input Input) (Output, error) {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "Generate JWT Token")
	defer span.End()

	err := uc.validator.Struct(input)
	if err != nil {
		return Output{}, err
	}

	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return Output{}, errs.ErrInvalidCredentials
		}
		return Output{}, err
	}

	if !user.IsActivated {
		return Output{}, errs.ErrUserIsNotActivated
	}

	hash := []byte(user.PasswordHash)
	pass := []byte(input.Password)
	err = uc.hashService.CompareHashAndPassword(hash, pass)
	if err != nil {
		return Output{}, errs.ErrInvalidCredentials
	}

	token, err := uc.generateTokenService.Execute(*user)
	if err != nil {
		return Output{}, err
	}

	return Output{Token: token}, nil
}
