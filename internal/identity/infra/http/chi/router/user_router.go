package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/chi/handler"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/middleware/auth_middleware"
)

func RegisterUserHandler(
	r *V1ChiRouter,
	userHandler *handler.UserHandler,
	authMiddleware *auth_middleware.Middleware,
) {
	r.Router.Post("/users", userHandler.Create)
	r.Router.Post("/users/activate", userHandler.Activate)
	r.Router.Get("/users/:id", authMiddleware.Execute, userHandler.FindByID)
	r.Router.Put("/users/:id", authMiddleware.Execute, userHandler.Update)
}
