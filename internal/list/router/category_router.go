package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/list/handler"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/middleware/auth_middleware"
)

func RegisterCategoryRoutes(
	r *V1FiberRouter,
	categoryHandler *handler.CategoryHandler,
	authMiddleware *auth_middleware.Middleware,
) {
	r.Post("/categories", authMiddleware.Execute, categoryHandler.Create)
	r.Post("/categories/find", authMiddleware.Execute, categoryHandler.Find)
	r.Put("/categories/:id", authMiddleware.Execute, categoryHandler.Update)
	r.Delete("/categories/:id", authMiddleware.Execute, categoryHandler.Delete)
}
