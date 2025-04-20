package service

import "github.com/cristiano-pacheco/shoplist/internal/identity/domain/model"

type EmailService interface {
	SendAccountConfirmationEmail(user model.UserModel, token string) error
}
