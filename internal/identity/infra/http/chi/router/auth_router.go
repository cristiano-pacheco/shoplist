package router

import "github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/chi/handler"

func RegisterAuthHandler(r *V1Router, authHandler *handler.AuthHandler) {
	r.Post("/auth/token", authHandler.GenerateToken)
}
