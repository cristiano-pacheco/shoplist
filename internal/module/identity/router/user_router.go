package router

import (
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/handler"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/middleware/auth_middleware"
)

func RegisterUserHandler(
	r *Router,
	userHandler *handler.UserHandler,
	authMiddleware *auth_middleware.Middleware,
) {
	r.Post("/users", userHandler.Store)
	r.Get("/users/:id", authMiddleware.Execute, userHandler.Show)
	r.Put("/users/:id", authMiddleware.Execute, userHandler.Update)
}
