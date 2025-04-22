package router

import (
	"github.com/cristiano-pacheco/shoplist/internal/kernel/http/httpserver"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	server *httpserver.HTTPServer
}

func NewRouter(server *httpserver.HTTPServer) *Router {
	return &Router{server: server}
}

func (r *Router) Router() *httprouter.Router {
	return r.server.Router()
}
