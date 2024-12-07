package auth_middleware

import (
	"context"
	"strconv"
	"strings"

	"github.com/cristiano-pacheco/shoplist/internal/shared/dto"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/response"
	"github.com/cristiano-pacheco/shoplist/internal/shared/registry"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	jwtParser          *jwt.Parser
	errorMapper        errs.ErrorMapper
	privateKeyRegistry registry.PrivateKeyRegistry
	isUserEnabledQuery *isUserEnabledQuery
}

func New(
	jwtParser *jwt.Parser,
	errorMapper errs.ErrorMapper,
	privateKeyRegistry registry.PrivateKeyRegistry,
	isUserEnabledQuery *isUserEnabledQuery,
) *Middleware {
	return &Middleware{jwtParser, errorMapper, privateKeyRegistry, isUserEnabledQuery}
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

	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return m.handleError(c, errs.ErrInvalidToken)
	}

	ctx := context.Background()
	isUserEnabled, err := m.isUserEnabledQuery.Execute(ctx, userID)
	if err != nil {
		return m.handleError(c, err)
	}

	if !isUserEnabled {
		return m.handleError(c, errs.ErrInvalidToken)
	}

	c.Locals("user_id", userID)

	return c.Next()
}

func (m *Middleware) handleError(c *fiber.Ctx, err error) error {
	rError := m.errorMapper.Map(err)
	return response.Error(c, rError)
}
