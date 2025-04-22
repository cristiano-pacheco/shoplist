package router

import (
	"net/http"

	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/handler"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/middleware"
)

func SetupUserRoutes(
	r *V1Router,
	userHandler *handler.UserHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	router := r.Router()
	router.HandlerFunc(http.MethodPost, "/users", userHandler.Create)
	router.HandlerFunc(http.MethodPost, "/users/activate", userHandler.Activate)
	router.HandlerFunc(http.MethodGet, "/users/:id", authMiddleware.Middleware(userHandler.FindByID))
	router.HandlerFunc(http.MethodPut, "/users/:id", authMiddleware.Middleware(userHandler.Update))
}
