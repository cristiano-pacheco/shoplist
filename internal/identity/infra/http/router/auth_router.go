package router

import (
	"net/http"

	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/handler"
)

func SetupAuthRoutes(r *V1Router, authHandler *handler.AuthHandler) {
	router := r.Router()
	router.HandlerFunc(http.MethodPost, "/auth/token", authHandler.GenerateToken)
}
