package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/kernel/http/httpserver"
	"github.com/julienschmidt/httprouter"
)

type V1Router struct {
	server *httpserver.HTTPServer
}

func NewV1Router(server *httpserver.HTTPServer) *V1Router {
	return &V1Router{server: server}
}

func (r *V1Router) Router() *httprouter.Router {
	return r.server.Router()
}
