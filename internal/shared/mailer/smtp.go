package mailer

import (
	"context"

	"github.com/go-mail/mail/v2"
	"go.opentelemetry.io/otel/trace"
)

type SmtpMailer interface {
	Send(ctx context.Context, md MailData) error
}

type smtpMailer struct {
	dialer *mail.Dialer
}

func NewSmtpMailer(dialer *mail.Dialer) SmtpMailer {
	return &smtpMailer{dialer}
}

func (m *smtpMailer) Send(ctx context.Context, md MailData) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	msg := mail.NewMessage()
	msg.SetHeader("To", md.ToEmail)
	msg.SetHeader("From", md.Sender)
	msg.SetHeader("Subject", md.Subject)
	msg.SetBody("text/html", md.Content)

	err := m.dialer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
