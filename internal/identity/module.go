package identity

import (
	"github.com/cristiano-pacheco/shoplist/internal/identity/application/usecase"
	domain_service "github.com/cristiano-pacheco/shoplist/internal/identity/domain/service"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/validator"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/handler"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/middleware"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/router"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/persistence/gorm/repository"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/service"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"identity",
	fx.Provide(
		// #################### APPLICATION ####################################
		// usecases
		usecase.NewUserCreateUseCase,
		usecase.NewUserActivateUseCase,
		usecase.NewUserUpdateUseCase,
		usecase.NewUserFindUseCase,
		usecase.NewTokenGenerateUseCase,

		// #################### DOMAIN #########################################
		domain_service.NewHashService,
		validator.NewPasswordValidator,

		// #################### INFRA ##########################################
		router.NewV1Router,

		// handlers
		handler.NewAuthHandler,
		handler.NewUserHandler,

		// middlewares
		middleware.NewAuthMiddleware,

		// mappers
		mapper.NewUserMapper,
		mapper.NewAccountConfirmationMapper,

		// repositories
		repository.NewUserRepository,
		repository.NewAccountConfirmationRepository,

		// services
		service.NewTokenService,
		service.NewEmailConfirmationService,
	),
	fx.Invoke(
		router.SetupUserRoutes,
		router.SetupAuthRoutes,
	),
)
