package main

import (
	"github.com/cristiano-pacheco/go-modulith/cmd"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
