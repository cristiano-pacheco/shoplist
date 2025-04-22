package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/service"
	domain_service "github.com/cristiano-pacheco/shoplist/internal/identity/domain/service"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/config"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/logger"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/mailer"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/otel"
)

const sendAccountConfirmationEmailTemplate = "account_confirmation.gohtml"
const sendAccountConfirmationEmailSubject = "Account Confirmation"

type SendEmailConfirmationService interface {
	service.SendEmailConfirmationService
}

type sendEmailConfirmationService struct {
	mailerTemplate                mailer.MailerTemplate
	mailer                        mailer.SmtpMailer
	accountConfirmationRepository repository.AccountConfirmationRepository
	userRepository                repository.UserRepository
	hashService                   domain_service.HashService
	logger                        logger.Logger
	cfg                           config.Config
}

func NewSendEmailConfirmationService(
	mailerTemplate mailer.MailerTemplate,
	smtpMailer mailer.SmtpMailer,
	accountConfirmationRepository repository.AccountConfirmationRepository,
	userRepository repository.UserRepository,
	hashService domain_service.HashService,
	logger logger.Logger,
	cfg config.Config,
) SendEmailConfirmationService {
	return &sendEmailConfirmationService{
		mailerTemplate,
		smtpMailer,
		accountConfirmationRepository,
		userRepository,
		hashService,
		logger,
		cfg,
	}
}

func (s *sendEmailConfirmationService) Execute(ctx context.Context, userID uint64) error {
	ctx, span := otel.Trace().StartSpan(ctx, "sendEmailConfirmationService.Execute")
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
		user.ID(),
		accountConfToken,
	)

	// compile the template
	tplData := struct {
		Name                    string
		AccountConfirmationLink string
	}{
		Name:                    user.Name(),
		AccountConfirmationLink: accountConfLink,
	}

	content, err := s.mailerTemplate.CompileTemplate(sendAccountConfirmationEmailTemplate, tplData)
	if err != nil {
		message := "error compiling template"
		s.logger.Error(message, "error", err)
		return err
	}

	acModel, err := model.CreateAccountConfirmationModel(user.ID(), accountConfToken, time.Now().Add(time.Hour*24))
	if err != nil {
		message := "error creating account confirmation model"
		s.logger.Error(message, "error", err)
		return err
	}

	// persist the account confirmation in the database
	_, err = s.accountConfirmationRepository.Create(ctx, acModel)
	if err != nil {
		message := "error creating account confirmation"
		s.logger.Error(message, "error", err)
		return err
	}

	md := mailer.MailData{
		Sender:  s.cfg.MAIL.Sender,
		ToName:  user.Name(),
		ToEmail: user.Email(),
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
