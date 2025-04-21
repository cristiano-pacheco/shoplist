package httpserver

import (
	"context"
	"fmt"

	_ "github.com/cristiano-pacheco/shoplist/docs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/middleware"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
)

type FiberServer struct {
	app  *fiber.App
	conf config.Config
}

func NewFiberHTTPServer(
	lc fx.Lifecycle,
	conf config.Config,
	errorHandlerMiddleware *middleware.ErrorHandlerMiddleware,
) *FiberServer {
	fiberConfig := fiber.Config{
		ReadBufferSize: 8192,
		ProxyHeader:    "X-Real-IP",
	}

	server := InitFiber(conf, errorHandlerMiddleware, fiberConfig)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			server.Run()
			return nil
		},
		OnStop: server.Shutdown,
	})

	return server
}

func InitFiber(
	conf config.Config,
	errorHandlerMiddleware *middleware.ErrorHandlerMiddleware,
	options ...fiber.Config,
) *FiberServer {
	var fiberConfig fiber.Config
	if len(options) > 0 {
		fiberConfig = options[0]
	}

	app := fiber.New(fiberConfig)

	app.Use(recover.New())
	app.Use(otelfiber.Middleware())
	app.Use(healthcheck.New())
	app.Use(errorHandlerMiddleware.Middleware())

	app.Get("/swagger/*", swagger.New())

	return &FiberServer{
		app:  app,
		conf: conf,
	}
}

func (s *FiberServer) Get(path string, handler ...fiber.Handler) {
	s.app.Get(path, handler...)
}

func (s *FiberServer) Post(path string, handler ...fiber.Handler) {
	s.app.Post(path, handler...)
}

func (s *FiberServer) Put(path string, handler ...fiber.Handler) {
	s.app.Put(path, handler...)
}

func (s *FiberServer) Patch(path string, handler ...fiber.Handler) {
	s.app.Patch(path, handler...)
}

func (s *FiberServer) Delete(path string, handler ...fiber.Handler) {
	s.app.Delete(path, handler...)
}

func (s *FiberServer) Group(path string, middleware ...fiber.Handler) fiber.Router {
	return s.app.Group(path, middleware...)
}

func (s *FiberServer) Run() {
	go func() {
		err := s.app.Listen(fmt.Sprintf(":%d", s.conf.HTTPPort))
		if err != nil {
			panic(err)
		}
	}()
}

func (s *FiberServer) App() *fiber.App {
	return s.app
}

func (s *FiberServer) Shutdown(context.Context) error {
	return s.app.Shutdown()
}

func (s *FiberServer) GetConfig() config.Config {
	return s.conf
}
