package repository

import "go.uber.org/fx"

var Module = fx.Module(
	"identity/repository",
	fx.Provide(NewUserRepository),
)
