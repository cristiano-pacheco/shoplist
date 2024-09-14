package generate_token_service

import (
	"encoding/base64"
	"strconv"
	"time"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/model"
	"github.com/golang-jwt/jwt/v5"
)

type ServiceI interface {
	Execute(user model.UserModel) (string, error)
}

type service struct {
	conf config.Config
}

func New(
	conf config.Config,
) ServiceI {
	return &service{conf}
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

	pkString, err := base64.StdEncoding.DecodeString(s.conf.JWT.PrivateKey)
	if err != nil {
		return "", err
	}

	privateKey, err := mapPEMToRSAPrivateKey(pkString)
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
