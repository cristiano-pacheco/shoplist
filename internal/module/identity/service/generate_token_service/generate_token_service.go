package generate_token_service

import (
	"strconv"
	"time"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/registry/privatekey_registry"
	"github.com/golang-jwt/jwt/v5"
)

type ServiceI interface {
	Execute(user model.UserModel) (string, error)
}

type service struct {
	conf               config.Config
	privateKeyRegistry privatekey_registry.RegistryI
}

func New(
	conf config.Config,
	privateKeyRegistry privatekey_registry.RegistryI,
) ServiceI {
	return &service{conf, privateKeyRegistry}
}

func (s *service) Execute(user model.UserModel) (string, error) {
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
		return "", err
	}

	return signedToken, nil
}
