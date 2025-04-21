package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/response"
	shared_jwt "github.com/cristiano-pacheco/shoplist/internal/shared/jwt"
	"github.com/cristiano-pacheco/shoplist/internal/shared/registry"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	jwtParser          *jwt.Parser
	errorMapper        errs.ErrorMapper
	privateKeyRegistry registry.PrivateKeyRegistry
	userRepository     repository.UserRepository
}

func NewAuthMiddleware(
	jwtParser *jwt.Parser,
	errorMapper errs.ErrorMapper,
	privateKeyRegistry registry.PrivateKeyRegistry,
	userRepository repository.UserRepository,
) *AuthMiddleware {
	return &AuthMiddleware{jwtParser, errorMapper, privateKeyRegistry, userRepository}
}

// UserIDKey is the key used to store the user ID in the request context
type contextKey string

const UserIDKey contextKey = "user_id"

// GetUserID extracts the user ID from the request context
func GetUserID(r *http.Request) (uint64, bool) {
	userID, ok := r.Context().Value(UserIDKey).(uint64)
	return userID, ok
}

// Middleware returns a Chi middleware function for authentication
func (m *AuthMiddleware) Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract and validate token
			bearerToken := r.Header.Get("Authorization")
			if !strings.HasPrefix(bearerToken, "Bearer ") {
				m.handleError(w, errs.ErrInvalidToken)
				return
			}

			jwtToken := strings.TrimSpace(bearerToken[7:])
			pk := m.privateKeyRegistry.Get()

			tokenKeyFunc := func(token *jwt.Token) (interface{}, error) {
				return &pk.PublicKey, nil
			}

			var claims shared_jwt.JWTClaims
			token, err := m.jwtParser.ParseWithClaims(jwtToken, &claims, tokenKeyFunc)
			if err != nil {
				m.handleError(w, errs.ErrInvalidToken)
				return
			}

			if !token.Valid {
				m.handleError(w, errs.ErrInvalidToken)
				return
			}

			userID, err := strconv.ParseUint(claims.Subject, 10, 64)
			if err != nil {
				m.handleError(w, errs.ErrInvalidToken)
				return
			}

			ctx := r.Context()
			isActivated, err := m.userRepository.IsActivated(ctx, userID)
			if err != nil {
				m.handleError(w, err)
				return
			}

			if !isActivated {
				m.handleError(w, errs.ErrUserIsNotActivated)
				return
			}

			// Store user ID in context
			ctx = context.WithValue(ctx, UserIDKey, userID)

			// Call next handler with updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (m *AuthMiddleware) handleError(w http.ResponseWriter, err error) {
	rError := m.errorMapper.Map(err)
	response.Error(w, rError)
}
