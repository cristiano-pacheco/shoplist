package usecase

import (
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/activate_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/create_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/find_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/generate_token_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/update_user_usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"identity/usecase",
	fx.Provide(
		create_user_usecase.New,
		activate_user_usecase.New,
		update_user_usecase.New,
		find_user_usecase.New,
		generate_token_usecase.New,
	),
)
