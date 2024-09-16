package auth_middleware

import (
	"strings"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/errs"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mapper/errormapper"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/registry/privatekey_registry"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	jwtParser          *jwt.Parser
	errorMapper        *errormapper.Mapper
	privateKeyRegistry privatekey_registry.RegistryI
}

func New(
	jwtParser *jwt.Parser,
	errorMapper *errormapper.Mapper,
	privateKeyRegistry privatekey_registry.RegistryI,
) *Middleware {
	return &Middleware{jwtParser, errorMapper, privateKeyRegistry}
}

func (m *Middleware) Execute(c *fiber.Ctx) error {
	bearerToken := c.Get("Authorization")
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return m.handleError(c, errs.ErrInvalidToken)
	}

	jwtToken := strings.TrimSpace(bearerToken[7:])
	pk := m.privateKeyRegistry.Get()

	tokenKeyFunc := func(token *jwt.Token) (interface{}, error) {
		return &pk.PublicKey, nil
	}

	var claims dto.JWTClaims
	token, err := m.jwtParser.ParseWithClaims(jwtToken, &claims, tokenKeyFunc)
	if err != nil {
		return m.handleError(c, errs.ErrInvalidToken)
	}

	if !token.Valid {
		return m.handleError(c, errs.ErrInvalidToken)
	}

	// TODO: check if the user is activated

	return c.Next()
}

func (m *Middleware) handleError(c *fiber.Ctx, err error) error {
	rError := m.errorMapper.MapErrorToResponseError(err)
	return response.HandleErrorResponse(c, rError)
}
