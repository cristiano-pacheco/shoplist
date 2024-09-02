package router

import "github.com/cristiano-pacheco/go-modulith/internal/module/identity/handler"

func RegisterUserHandler(r *Router, userHandler *handler.UserHandler) {
	r.Post("/users", userHandler.Store)
	r.Get("/users/:id", userHandler.Show)
	r.Put("/users/:id", userHandler.Update)
}
