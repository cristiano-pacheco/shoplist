package mailer

import "go.uber.org/fx"

var Module = fx.Module(
	"shared/mailer",
	fx.Provide(NewMailerTemplate),
	fx.Provide(NewSmtpMailer),
	fx.Provide(NewDialer),
)
