package service

import (
	generate_token_service "github.com/cristiano-pacheco/go-modulith/internal/identity/service/generate_token"
	hash_service "github.com/cristiano-pacheco/go-modulith/internal/identity/service/hash"
	send_account_confirmation_email_service "github.com/cristiano-pacheco/go-modulith/internal/identity/service/send_account_confirmation_email"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"identity/service",
	fx.Provide(generate_token_service.New),
	fx.Provide(hash_service.New),
	fx.Provide(send_account_confirmation_email_service.New),
)
