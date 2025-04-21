package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/chi/handler"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/chi/middleware"
)

func SetupUserRoutes(
	r *V1ChiRouter,
	userHandler *handler.UserHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	r.Router.Post("/users", userHandler.Create)
	r.Router.Post("/users/activate", userHandler.Activate)

	r.Router.With(authMiddleware.Middleware()).Get("/users/{id}", userHandler.FindByID)
	r.Router.With(authMiddleware.Middleware()).Put("/users/{id}", userHandler.Update)
}
