package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "github.com/cristiano-pacheco/shoplist/docs"
	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type ChiHTTPServer struct {
	router chi.Router
	server *http.Server
}

func NewChiHTTPServer(
	corsConfig CorsConfig,
	otelHandlerName string,
	isOtelEnabled bool,
	httpPort uint,
) *ChiHTTPServer {
	r := chi.NewRouter()

	// Basic middleware
	r.Use(chi_middleware.Logger)
	r.Use(chi_middleware.Recoverer)
	r.Use(chi_middleware.RequestID)
	r.Use(chi_middleware.RealIP)
	r.Use(chi_middleware.Timeout(30 * time.Second))

	corsOptions := cors.Options{
		AllowedOrigins:   corsConfig.AllowedOrigins,
		AllowedMethods:   corsConfig.AllowedMethods,
		AllowedHeaders:   corsConfig.AllowedHeaders,
		ExposedHeaders:   corsConfig.ExposedHeaders,
		AllowCredentials: corsConfig.AllowCredentials,
		MaxAge:           corsConfig.MaxAge,
	}

	// CORS configuration
	r.Use(cors.Handler(corsOptions))

	if isOtelEnabled {
		// OpenTelemetry middleware
		r.Use(func(next http.Handler) http.Handler {
			return otelhttp.NewHandler(next, otelHandlerName)
		})
	}

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
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", httpPort),
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

func (s *ChiHTTPServer) Router() chi.Router {
	return s.router
}

func (s *ChiHTTPServer) Run() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *ChiHTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
