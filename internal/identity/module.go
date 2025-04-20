package identity

import (
	"github.com/cristiano-pacheco/shoplist/internal/identity/application/usecase"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/chi/handler"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/chi/router"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/repository"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/service"
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
