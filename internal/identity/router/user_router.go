package router

import (
	"github.com/cristiano-pacheco/go-modulith/internal/identity/handler"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/http/middleware/auth_middleware"
)

func RegisterUserHandler(
	r *Router,
	userHandler *handler.UserHandler,
	authMiddleware *auth_middleware.Middleware,
) {
	r.Post("/users", userHandler.Create)
	r.Post("/users/activate", userHandler.Activate)
	r.Get("/users/:id", authMiddleware.Execute, userHandler.FindByID)
	r.Put("/users/:id", authMiddleware.Execute, userHandler.Update)
}
