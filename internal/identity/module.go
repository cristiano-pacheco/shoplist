package identity

import (
	"github.com/cristiano-pacheco/shoplist/internal/identity/application/usecase"
	domain_repository "github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
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
		router.NewRouter,

		// handlers
		handler.NewAuthHandler,
		handler.NewUserHandler,

		// middlewares
		middleware.NewAuthMiddleware,

		// mappers
		mapper.NewUserMapper,
		mapper.NewAccountConfirmationMapper,

		// repositories
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(domain_repository.UserRepository)),
		),

		fx.Annotate(
			repository.NewAccountConfirmationRepository,
			fx.As(new(domain_repository.AccountConfirmationRepository)),
		),

		// services
		fx.Annotate(
			service.NewSendEmailConfirmationService,
			fx.As(new(domain_service.SendEmailConfirmationService)),
		),

		fx.Annotate(
			service.NewTokenService,
			fx.As(new(domain_service.TokenService)),
		),
	),
	fx.Invoke(
		router.SetupUserRoutes,
		router.SetupAuthRoutes,
	),
)
