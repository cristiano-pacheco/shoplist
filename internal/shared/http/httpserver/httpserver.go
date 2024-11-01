package httpserver

import (
	"context"
	"fmt"
	"time"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/fx"
)

type Server struct {
	app  *fiber.App
	conf config.Config
}

func NewHTTPServer(lc fx.Lifecycle, conf config.Config) *Server {
	fiberConfig := fiber.Config{
		ReadBufferSize: 8192,
		ProxyHeader:    "X-Real-IP",
	}

	server := Init(conf, fiberConfig)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			server.Run()
			return nil
		},
		OnStop: server.Shutdown,
	})

	return server
}

func Init(conf config.Config, options ...fiber.Config) *Server {
	config := fiber.Config{
		EnablePrintRoutes: !conf.IsProduction(),
		AppName:           conf.App.Name,
		IdleTimeout:       5 * time.Second,
	}

	app := fiber.New(config)
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(healthcheck.New())

	return &Server{app, conf}
}

func (s *Server) Get(path string, handler ...fiber.Handler) {
	s.app.Get(path, handler...)
}

func (s *Server) Post(path string, handler ...fiber.Handler) {
	s.app.Post(path, handler...)
}

func (s *Server) Put(path string, handler ...fiber.Handler) {
	s.app.Put(path, handler...)
}

func (s *Server) Patch(path string, handler ...fiber.Handler) {
	s.app.Patch(path, handler...)
}

func (s *Server) Delete(path string, handler ...fiber.Handler) {
	s.app.Delete(path, handler...)
}

func (s *Server) Group(path string, middleware ...fiber.Handler) fiber.Router {
	return s.app.Group(path, middleware...)
}

func (s *Server) Run() {
	go func() {
		err := s.app.Listen(fmt.Sprintf(":%d", s.conf.HTTPPort))
		if err != nil {
			panic(err)
		}
	}()
}

func (s *Server) App() *fiber.App {
	return s.app
}

func (s *Server) Shutdown(context.Context) error {
	return s.app.Shutdown()
}
