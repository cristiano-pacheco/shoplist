package identity

import (
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/handler"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/router"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"identity",
	fx.Provide(
		router.NewRouter,

		// handlers
		handler.NewAuthHandler,
		handler.NewUserHandler,

		// repositories
		repository.NewUserRepository,
		repository.NewAccountConfirmationRepository,

		// services
		service.NewTokenService,
		service.NewHashService,
		service.NewEmailConfirmationService,

		// usecases
		usecase.NewUserCreateUseCase,
		usecase.NewUserActivateUseCase,
		usecase.NewUserUpdateUseCase,
		usecase.NewUserFindUseCase,
		usecase.NewTokenGenerateUseCase,
	),
	fx.Invoke(
		router.RegisterUserHandler,
		router.RegisterAuthHandler,
	),
)
