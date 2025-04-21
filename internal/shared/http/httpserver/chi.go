package httpserver

import (
	"context"
	"net/http"

	_ "github.com/cristiano-pacheco/shoplist/docs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/middleware"
	"github.com/cristiano-pacheco/shoplist/pkg/httpserver"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type ChiHTTPServer struct {
	server *httpserver.ChiHTTPServer
}

func NewChiHTTPServer(
	lc fx.Lifecycle,
	conf config.Config,
	errorHandlerMiddleware *middleware.ErrorHandlerMiddleware,
) *ChiHTTPServer {
	corsConfig := httpserver.CorsConfig{
		AllowedOrigins:   conf.CORS.GetAllowedOrigins(),
		AllowedMethods:   conf.CORS.GetAllowedMethods(),
		AllowedHeaders:   conf.CORS.GetAllowedHeaders(),
		ExposedHeaders:   conf.CORS.GetExposedHeaders(),
		AllowCredentials: conf.CORS.AllowCredentials,
		MaxAge:           conf.CORS.MaxAge,
	}

	isOtelEnabled := true
	server := httpserver.NewChiHTTPServer(corsConfig, conf.App.Name, isOtelEnabled, conf.HTTPPort)
	router := server.Router()
	router.Use(errorHandlerMiddleware.Middleware())

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			server.Run()
			return nil
		},
		OnStop: server.Shutdown,
	})

	return &ChiHTTPServer{server}
}

func (s *ChiHTTPServer) Get(path string, handler http.HandlerFunc) {
	s.server.Get(path, handler)
}

func (s *ChiHTTPServer) Post(path string, handler http.HandlerFunc) {
	s.server.Post(path, handler)
}

func (s *ChiHTTPServer) Put(path string, handler http.HandlerFunc) {
	s.server.Put(path, handler)
}

func (s *ChiHTTPServer) Patch(path string, handler http.HandlerFunc) {
	s.server.Patch(path, handler)
}

func (s *ChiHTTPServer) Delete(path string, handler http.HandlerFunc) {
	s.server.Delete(path, handler)
}

func (s *ChiHTTPServer) Group(path string) chi.Router {
	return s.server.Group(path)
}

func (s *ChiHTTPServer) Router() chi.Router {
	return s.server.Router()
}
