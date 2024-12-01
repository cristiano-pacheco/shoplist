package service

import (
	"github.com/cristiano-pacheco/go-modulith/internal/modules/identity/service/generate_token_service"
	"github.com/cristiano-pacheco/go-modulith/internal/modules/identity/service/hash_service"
	"github.com/cristiano-pacheco/go-modulith/internal/modules/identity/service/send_account_confirmation_email_service"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"identity/service",
	fx.Provide(
		generate_token_service.New,
		hash_service.New,
		send_account_confirmation_email_service.New,
	),
)
