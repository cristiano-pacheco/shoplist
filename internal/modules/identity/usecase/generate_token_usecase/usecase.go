package generate_token_usecase

import (
	"context"
	"errors"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service/generate_token_service"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service/hash_service"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/telemetry"
	"github.com/cristiano-pacheco/shoplist/internal/shared/validator"
)

type UseCase struct {
	validator            validator.ValidateI
	userRepo             repository.UserRepositoryI
	hashService          hash_service.ServiceI
	generateTokenService generate_token_service.ServiceI
}

func New(
	validator validator.ValidateI,
	userRepo repository.UserRepositoryI,
	hashService hash_service.ServiceI,
	generateTokenService generate_token_service.ServiceI,
) *UseCase {
	return &UseCase{
		validator,
		userRepo,
		hashService,
		generateTokenService,
	}
}

func (uc *UseCase) Execute(ctx context.Context, input Input) (Output, error) {
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
