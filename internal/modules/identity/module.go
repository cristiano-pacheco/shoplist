package identity

import (
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/handler"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/router"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/activate_user"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/create_user"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/find_user"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/generate_token"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/update_user"
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
		create_user.New,
		activate_user.New,
		update_user.New,
		find_user.New,
		generate_token.New,
	),
	fx.Invoke(
		router.RegisterUserHandler,
		router.RegisterAuthHandler,
	),
)
