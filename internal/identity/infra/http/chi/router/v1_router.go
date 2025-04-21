package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/kernel/http/httpserver"
	"github.com/go-chi/chi/v5"
)

type V1ChiRouter struct {
	Router chi.Router
	Server *httpserver.ChiHTTPServer
}

func NewChiV1Router(server *httpserver.ChiHTTPServer) *V1ChiRouter {
	router := server.Group("/api/v1")
	return &V1ChiRouter{Router: router, Server: server}
}
