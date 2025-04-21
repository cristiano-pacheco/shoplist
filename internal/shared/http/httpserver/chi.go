package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "github.com/cristiano-pacheco/shoplist/docs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/middleware"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/fx"
)

type Server struct {
	router chi.Router
	conf   config.Config
	server *http.Server
}

func NewHTTPServer(
	lc fx.Lifecycle,
	conf config.Config,
	errorHandlerMiddleware *middleware.ErrorHandlerMiddleware,
) *Server {
	server := Init(conf, errorHandlerMiddleware)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			server.Run()
			return nil
		},
		OnStop: server.Shutdown,
	})

	return server
}

func Init(
	conf config.Config,
	errorHandlerMiddleware *middleware.ErrorHandlerMiddleware,
) *Server {
	r := chi.NewRouter()

	// Basic middleware
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Timeout(30 * time.Second))

	// CORS configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// OpenTelemetry middleware
	r.Use(func(next http.Handler) http.Handler {
		return otelhttp.NewHandler(next, "http-server")
	})

	// Error handler middleware will be implemented as needed
	// We'll need to adapt it from Fiber to Chi

	// Health check
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Metrics endpoint
	r.Handle("/metrics", promhttp.Handler())

	// Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &Server{
		router: r,
		conf:   conf,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", conf.HTTPPort),
			Handler: r,
		},
	}

	return server
}

func (s *Server) Get(path string, handler http.HandlerFunc) {
	s.router.Get(path, handler)
}

func (s *Server) Post(path string, handler http.HandlerFunc) {
	s.router.Post(path, handler)
}

func (s *Server) Put(path string, handler http.HandlerFunc) {
	s.router.Put(path, handler)
}

func (s *Server) Patch(path string, handler http.HandlerFunc) {
	s.router.Patch(path, handler)
}

func (s *Server) Delete(path string, handler http.HandlerFunc) {
	s.router.Delete(path, handler)
}

func (s *Server) Group(path string) chi.Router {
	return s.router.Route(path, nil)
}

func (s *Server) Run() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *Server) Router() chi.Router {
	return s.router
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) GetConfig() config.Config {
	return s.conf
}
