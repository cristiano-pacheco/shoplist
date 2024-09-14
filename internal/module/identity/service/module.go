package service

import (
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/service/generate_token_service"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/service/hash_service"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"identity/service",
	fx.Provide(generate_token_service.New),
	fx.Provide(hash_service.New),
)
