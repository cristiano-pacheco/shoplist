package main

import (
	"github.com/cristiano-pacheco/go-modulith/cmd"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
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
	telemetry.Init(cfg)
	cmd.Execute()
}
