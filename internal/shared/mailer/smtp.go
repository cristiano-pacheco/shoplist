package mailer

import (
	"github.com/go-mail/mail/v2"
)

type SmtpMailer interface {
	Send(md MailData) error
}

type smtpMailer struct {
	dialer *mail.Dialer
}

func NewSmtpMailer(dialer *mail.Dialer) SmtpMailer {
	return &smtpMailer{dialer}
}

func (m *smtpMailer) Send(md MailData) error {
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
