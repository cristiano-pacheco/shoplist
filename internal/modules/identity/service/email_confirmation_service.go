package service

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/model"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/mailer"
	"go.opentelemetry.io/otel/trace"
)

const sendAccountConfirmationEmailTemplate = "account_confirmation.gohtml"
const sendAccountConfirmationEmailSubject = "Account Confirmation"

type EmailConfirmationService interface {
	Send(ctx context.Context, userID uint64) error
}

type emailConfirmationService struct {
	mailerTemplate                mailer.MailerTemplate
	mailer                        mailer.SmtpMailer
	accountConfirmationRepository repository.AccountConfirmationRepository
	userRepository                repository.UserRepository
	hashService                   HashService
	logger                        logger.Logger
	cfg                           config.Config
}

func NewEmailConfirmationService(
	mailerTemplate mailer.MailerTemplate,
	smtpMailer mailer.SmtpMailer,
	accountConfirmationRepository repository.AccountConfirmationRepository,
	userRepository repository.UserRepository,
	hashService HashService,
	logger logger.Logger,
	cfg config.Config,
) EmailConfirmationService {
	return &emailConfirmationService{
		mailerTemplate,
		smtpMailer,
		accountConfirmationRepository,
		userRepository,
		hashService,
		logger,
		cfg,
	}
}

func (s *emailConfirmationService) Send(ctx context.Context, userID uint64) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		message := "error finding user"
		s.logger.Error(message, "error", err)
		return err
	}
	// generate a random token
	token, err := s.hashService.GenerateRandomBytes()
	if err != nil {
		message := "error generating random bytes"
		s.logger.Error(message, "error", err)
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

	content, err := s.mailerTemplate.CompileTemplate(sendAccountConfirmationEmailTemplate, tplData)
	if err != nil {
		message := "error compiling template"
		s.logger.Error(message, "error", err)
		return err
	}

	// create the account confirmation model
	acModel := model.AccountConfirmationModel{
		UserID: user.ID,
		Token:  accountConfToken,
	}

	// persist the account confirmation in the database
	err = s.accountConfirmationRepository.Create(ctx, acModel)
	if err != nil {
		message := "error creating account confirmation"
		s.logger.Error(message, "error", err)
		return err
	}

	md := mailer.MailData{
		Sender:  s.cfg.MAIL.Sender,
		ToName:  user.Name,
		ToEmail: user.Email,
		Subject: sendAccountConfirmationEmailSubject,
		Content: content,
	}

	err = s.mailer.Send(ctx, md)
	if err != nil {
		message := "error sending email"
		s.logger.Error(message, "error", err)
		return err
	}

	return nil
}
