package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/httpserver"
	"github.com/gofiber/fiber/v2"
)

type V1Router struct {
	fiber.Router
	Server *httpserver.Server
}

func NewRouter(server *httpserver.Server) *V1Router {
	router := server.Group("/api/v1")
	return &V1Router{Router: router, Server: server}
}
