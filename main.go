package main

import (
	"context"
	"log"

	"github.com/cristiano-pacheco/go-modulith/cmd"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
)

func main() {
	config.Init()
	telemetry.Init()
	defer func() {
		t := telemetry.Get()
		if err := t.Shutdown(context.Background()); err != nil {
			log.Println(err)
		}
	}()
	cmd.Execute()
}
