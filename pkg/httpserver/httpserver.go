package httpserver

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/cristiano-pacheco/shoplist/docs"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type HTTPServer struct {
	router *httprouter.Router
	server *http.Server
}

func NewHTTPServer(
	corsConfig CorsConfig,
	otelHandlerName string,
	isOtelEnabled bool,
	httpPort uint,
) *HTTPServer {
	r := httprouter.New()

	// CORS configuration
	if len(corsConfig.AllowedOrigins) > 0 {
		r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Access-Control-Request-Method") != "" {
				// Set CORS headers
				header := w.Header()
				header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
				header.Set("Access-Control-Allow-Origin", corsConfig.AllowedOrigins[0])
				if len(corsConfig.AllowedHeaders) > 0 {
					header.Set("Access-Control-Allow-Headers", join(corsConfig.AllowedHeaders, ", "))
				}
				if corsConfig.AllowCredentials {
					header.Set("Access-Control-Allow-Credentials", "true")
				}
				if corsConfig.MaxAge > 0 {
					header.Set("Access-Control-Max-Age", fmt.Sprintf("%d", corsConfig.MaxAge))
				}
			}
			w.WriteHeader(http.StatusNoContent)
		})
	}

	// Health check
	r.GET("/healthcheck", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Metrics endpoint
	r.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	// Swagger
	r.Handler(http.MethodGet, "/swagger/*filepath", httpSwagger.WrapHandler)

	server := &HTTPServer{
		router: r,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", httpPort),
			Handler: r,
		},
	}

	// Apply OpenTelemetry if enabled
	if isOtelEnabled {
		server.server.Handler = otelhttp.NewHandler(r, otelHandlerName)
	}

	return server
}

func (s *HTTPServer) Router() *httprouter.Router {
	return s.router
}

func (s *HTTPServer) Run() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func join(s []string, sep string) string {
	if len(s) == 0 {
		return ""
	}
	result := s[0]
	for i := 1; i < len(s); i++ {
		result += sep + s[i]
	}
	return result
}
