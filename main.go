package main

import (
	"github.com/cristiano-pacheco/go-modulith/cmd"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
)

func main() {
	config.Init()
	cfg := config.GetConfig()
	telemetry.Init(cfg)
	cmd.Execute()
}
