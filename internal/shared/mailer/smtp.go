package mailer

import (
	"github.com/go-mail/mail/v2"
)

type SmtpMailerI interface {
	Send(md MailData) error
}

type smtpMailer struct {
	dialer *mail.Dialer
	sender string
}

func NewSmtpMailer(dialer *mail.Dialer, sender string) SmtpMailerI {
	return &smtpMailer{dialer, sender}
}

func (m *smtpMailer) Send(md MailData) error {
	msg := mail.NewMessage()
	msg.SetHeader("To", md.ToEmail)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", md.Subject)
	msg.SetBody("text/html", md.Content)

	err := m.dialer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
