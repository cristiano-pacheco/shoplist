package send_account_confirmation_email_service

import (
	"context"
	"encoding/base64"
	"fmt"

	identity_model "github.com/cristiano-pacheco/go-modulith/internal/module/identity/model"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/repository"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/service/hash_service"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mailer"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
)

const emailTemplate = "account_confirmation.gohtml"
const emailSubject = "Account Confirmation"

type ServiceI interface {
	Execute(ctx context.Context, user model.UserModel) error
}

type service struct {
	mailerTemplate                mailer.MailerTemplateI
	mailer                        mailer.SmtpMailerI
	accountConfirmationRepository repository.AccountConfirmationRepositoryI
	hashService                   hash_service.ServiceI
	cfg                           config.Config
}

func New(
	mailerTemplate mailer.MailerTemplateI,
	smtpMailer mailer.SmtpMailerI,
	accountConfirmationRepository repository.AccountConfirmationRepositoryI,
	hashService hash_service.ServiceI,
	cfg config.Config,
) ServiceI {
	return &service{mailerTemplate, smtpMailer, accountConfirmationRepository, hashService, cfg}
}

func (s *service) Execute(ctx context.Context, user model.UserModel) error {
	t := telemetry.Get()
	ctx, span := t.StartSpan(ctx, "send_account_confirmation_email_service.execute")
	defer span.End()

	// generate a random token
	token, err := s.hashService.GenerateRandomBytes()
	if err != nil {
		return err
	}

	// encode the token
	accountConfToken := base64.StdEncoding.EncodeToString(token)

	// generate the account confirmation link
	accountConfLink := fmt.Sprintf(
		"%s/user/confirmation?id=%d&token=%s",
		s.cfg.App.BaseURL,
		user.ID,
		accountConfToken,
	)

	// compile the template
	tplData := struct {
		Name                    string
		AccountConfirmationLink string
	}{
		Name:                    user.Name,
		AccountConfirmationLink: accountConfLink,
	}

	content, err := s.mailerTemplate.CompileTemplate(emailTemplate, tplData)
	if err != nil {
		return err
	}

	// create the account confirmation model
	acModel := identity_model.AccountConfirmationModel{
		UserID: user.ID,
		Token:  accountConfToken,
	}

	// persist the account confirmation in the database
	err = s.accountConfirmationRepository.Create(ctx, acModel)
	if err != nil {
		return err
	}

	md := mailer.MailData{
		Sender:  s.cfg.MAIL.Sender,
		ToName:  user.Name,
		ToEmail: user.Email,
		Subject: emailSubject,
		Content: content,
	}

	err = s.mailer.Send(md)
	if err != nil {
		return err
	}

	return nil
}
