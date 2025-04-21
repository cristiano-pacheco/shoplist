package router

import "github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/chi/handler"

func RegisterAuthHandler(r *V1ChiRouter, authHandler *handler.AuthHandler) {
	r.Router.Post("/auth/token", authHandler.GenerateToken)
}
