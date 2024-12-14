package main

import (
	"context"
	"log"

	"github.com/cristiano-pacheco/shoplist/cmd"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/pkg/trace"
)

// @title           Go modulith API
// @version         1.0
// @description     Go modulith API

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer <token>**

// @BasePath  /
func main() {
	config.Init()
	//cfg := config.GetConfig()
	//telemetry.Init(cfg)

	cfg := config.GetConfig()
	tracerCfg := trace.TracerConfig{
		ServiceName: cfg.App.Name,
		TracerURL:   cfg.Telemetry.TracerURL,
	}

	tp := trace.InitTracer(tracerCfg)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	cmd.Execute()
}
