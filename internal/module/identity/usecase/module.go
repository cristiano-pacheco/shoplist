package usecase

import (
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/create_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/find_activated_user_by_id_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/find_user_by_id_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/generate_jwt_token_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/update_user_usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"identity/usecase",
	fx.Provide(create_user_usecase.New),
	fx.Provide(find_user_by_id_usecase.New),
	fx.Provide(update_user_usecase.New),
	fx.Provide(generate_jwt_token_usecase.New),
	fx.Provide(find_activated_user_by_id_usecase.New),
)
