package router

import "go.uber.org/fx"

var Module = fx.Module("identity/router",
	fx.Provide(NewRouter),
	fx.Invoke(RegisterUserHandler),
	fx.Invoke(RegisterAuthHandler),
)
