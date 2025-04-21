package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/list/handler"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/middleware/auth_middleware"
)

func RegisterListRoutes(
	r *V1Router,
	listHandler *handler.ListHandler,
	authMiddleware *auth_middleware.Middleware,
) {
	r.Post("/lists", authMiddleware.Execute, listHandler.Create)
	r.Get("/lists", authMiddleware.Execute, listHandler.FindByUserID)
	r.Get("/lists/:id", authMiddleware.Execute, listHandler.FindByID)
	r.Put("/lists/:id", authMiddleware.Execute, listHandler.Update)
	r.Delete("/lists/:id", authMiddleware.Execute, listHandler.Delete)
}
