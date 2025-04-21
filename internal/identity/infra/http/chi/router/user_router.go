package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/chi/handler"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/middleware/auth_middleware"
)

func RegisterUserHandler(
	r *V1FiberRouter,
	userHandler *handler.UserHandler,
	authMiddleware *auth_middleware.Middleware,
) {
	r.Post("/users", userHandler.Create)
	r.Post("/users/activate", userHandler.Activate)
	r.Get("/users/:id", authMiddleware.Execute, userHandler.FindByID)
	r.Put("/users/:id", authMiddleware.Execute, userHandler.Update)
}
