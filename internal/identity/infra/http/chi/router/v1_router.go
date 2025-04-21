package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/httpserver"
	"github.com/gofiber/fiber/v2"
)

type V1FiberRouter struct {
	fiber.Router
	Server *httpserver.FiberServer
}

func NewV1FiberRouter(server *httpserver.FiberServer) *V1FiberRouter {
	router := server.Group("/api/v1")
	return &V1FiberRouter{Router: router, Server: server}
}
