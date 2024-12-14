package main

import (
	"context"
	"log"

	"github.com/cristiano-pacheco/shoplist/cmd"
	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
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

	cfg := config.GetConfig()
	otel.Init(cfg)

	defer func() {
		if err := otel.Trace().Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	cmd.Execute()
}
