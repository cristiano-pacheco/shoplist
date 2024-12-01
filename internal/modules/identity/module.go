package identity

import (
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/handler"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/router"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service/generate_token_service"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service/hash_service"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/service/send_account_confirmation_email_service"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/activate_user_usecase"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/create_user_usecase"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/find_user_usecase"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/generate_token_usecase"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/update_user_usecase"
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
		generate_token_service.New,
		hash_service.New,
		send_account_confirmation_email_service.New,

		// usecases
		create_user_usecase.New,
		activate_user_usecase.New,
		update_user_usecase.New,
		find_user_usecase.New,
		generate_token_usecase.New,
	),
	fx.Invoke(
		router.RegisterUserHandler,
		router.RegisterAuthHandler,
	),
)
