package router

import (
	"github.com/cristiano-pacheco/go-modulith/internal/shared/http/httpserver"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	fiber.Router
	Server *httpserver.Server
}

func NewRouter(server *httpserver.Server) *Router {
	router := server.Group("/api/v1")
	return &Router{Router: router, Server: server}
}
