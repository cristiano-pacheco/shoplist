package main

import (
	"github.com/cristiano-pacheco/go-modulith/cmd"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
)

func main() {
	config.Init()
	telemetry.Init()
	cmd.Execute()
}
