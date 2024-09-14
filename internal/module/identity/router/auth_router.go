package router

import "github.com/cristiano-pacheco/go-modulith/internal/module/identity/handler"

func RegisterAuthHandler(r *Router, authHandler *handler.AuthHandler) {
	r.Post("/auth/token", authHandler.Execute)
}
