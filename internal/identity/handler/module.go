package handler

import "go.uber.org/fx"

var Module = fx.Module(
	"identity/handler",
	fx.Provide(NewUserHandler),
	fx.Provide(NewAuthHandler),
)
