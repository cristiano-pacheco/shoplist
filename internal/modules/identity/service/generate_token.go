package service

import (
	"strconv"
	"time"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/model"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/logger"
	"github.com/cristiano-pacheco/shoplist/internal/shared/registry/privatekey_registry"
	"github.com/golang-jwt/jwt/v5"
)

type GenerateTokenServiceI interface {
	Execute(user model.UserModel) (string, error)
}

type GenerateTokenService struct {
	privateKeyRegistry privatekey_registry.RegistryI
	conf               config.Config
	logger             logger.LoggerI
}

func NewGenerateTokenService(
	conf config.Config,
	privateKeyRegistry privatekey_registry.RegistryI,
	logger logger.LoggerI,
) GenerateTokenServiceI {
	return &GenerateTokenService{privateKeyRegistry, conf, logger}
}

func (s *GenerateTokenService) Execute(user model.UserModel) (string, error) {
	now := time.Now()
	duration := time.Duration(s.conf.JWT.ExpirationInSeconds) * time.Second
	expires := now.Add(duration)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expires),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		Issuer:    s.conf.JWT.Issuer,
		Subject:   strconv.FormatUint(user.ID, 10),
	}

	method := jwt.GetSigningMethod(jwt.SigningMethodRS256.Name)
	token := jwt.NewWithClaims(method, claims)

	pk := s.privateKeyRegistry.Get()
	signedToken, err := token.SignedString(pk)
	if err != nil {
		message := "[generate_token] error signing token"
		s.logger.Error(message, "error", err)
		return "", err
	}

	return signedToken, nil
}
