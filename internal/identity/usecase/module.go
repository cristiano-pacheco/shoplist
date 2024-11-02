package usecase

import (
	create_user_usecase "github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/create_user"
	find_user_by_id_usecase "github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/find_user_by_id"
	generate_jwt_token_usecase "github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/generate_jwt_token"
	update_user_usecase "github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/update_user"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"identity/usecase",
	fx.Provide(create_user_usecase.New),
	fx.Provide(find_user_by_id_usecase.New),
	fx.Provide(update_user_usecase.New),
	fx.Provide(generate_jwt_token_usecase.New),
)
