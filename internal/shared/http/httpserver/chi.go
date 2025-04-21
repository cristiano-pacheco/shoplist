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
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/fx"
)

type ChiHTTPServer struct {
	router chi.Router
	conf   config.Config
	server *http.Server
}

func NewChiHTTPServer(
	lc fx.Lifecycle,
	conf config.Config,
	errorHandlerMiddleware *middleware.ErrorHandlerMiddleware,
) *ChiHTTPServer {
	server := InitChiHTTPServer(conf, errorHandlerMiddleware)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			server.Run()
			return nil
		},
		OnStop: server.Shutdown,
	})

	return server
}

func InitChiHTTPServer(
	conf config.Config,
	errorHandlerMiddleware *middleware.ErrorHandlerMiddleware,
) *ChiHTTPServer {
	r := chi.NewRouter()

	// Basic middleware
	r.Use(chi_middleware.Logger)
	r.Use(chi_middleware.Recoverer)
	r.Use(chi_middleware.RequestID)
	r.Use(chi_middleware.RealIP)
	r.Use(chi_middleware.Timeout(30 * time.Second))

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
		return otelhttp.NewHandler(next, "chi-http-server")
	})

	// Error handler middleware
	r.Use(errorHandlerMiddleware.ChiMiddleware())

	// Health check
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Metrics endpoint
	r.Handle("/metrics", promhttp.Handler())

	// Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &ChiHTTPServer{
		router: r,
		conf:   conf,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", conf.HTTPPort),
			Handler: r,
		},
	}

	return server
}

func (s *ChiHTTPServer) Get(path string, handler http.HandlerFunc) {
	s.router.Get(path, handler)
}

func (s *ChiHTTPServer) Post(path string, handler http.HandlerFunc) {
	s.router.Post(path, handler)
}

func (s *ChiHTTPServer) Put(path string, handler http.HandlerFunc) {
	s.router.Put(path, handler)
}

func (s *ChiHTTPServer) Patch(path string, handler http.HandlerFunc) {
	s.router.Patch(path, handler)
}

func (s *ChiHTTPServer) Delete(path string, handler http.HandlerFunc) {
	s.router.Delete(path, handler)
}

func (s *ChiHTTPServer) Group(path string) chi.Router {
	return s.router.Route(path, nil)
}

func (s *ChiHTTPServer) Run() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *ChiHTTPServer) Router() chi.Router {
	return s.router
}

func (s *ChiHTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
