package httpserver

import (
	"context"

	_ "github.com/cristiano-pacheco/shoplist/docs"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/config"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/http/middleware"
	"github.com/cristiano-pacheco/shoplist/pkg/httpserver"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/fx"
)

type HTTPServer struct {
	server *httpserver.HTTPRouterServer
}

func NewHTTPServer(
	lc fx.Lifecycle,
	conf config.Config,
	errorHandlerMiddleware *middleware.ErrorHandlerMiddleware,
) *HTTPServer {
	corsConfig := httpserver.CorsConfig{
		AllowedOrigins:   conf.CORS.GetAllowedOrigins(),
		AllowedMethods:   conf.CORS.GetAllowedMethods(),
		AllowedHeaders:   conf.CORS.GetAllowedHeaders(),
		ExposedHeaders:   conf.CORS.GetExposedHeaders(),
		AllowCredentials: conf.CORS.AllowCredentials,
		MaxAge:           conf.CORS.MaxAge,
	}

	isOtelEnabled := true
	server := httpserver.NewHTTPRouterServer(corsConfig, conf.App.Name, isOtelEnabled, conf.HTTPPort)

	httpRouterServer := &HTTPServer{
		server: server,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			server.Run()
			return nil
		},
		OnStop: server.Shutdown,
	})

	return httpRouterServer
}

func (s *HTTPServer) Router() *httprouter.Router {
	return s.server.Router()
}
