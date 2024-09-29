package telemetry

import (
	"log"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	pkg_telemetry "github.com/cristiano-pacheco/go-modulith/pkg/telemetry"
)

var _global *pkg_telemetry.Telemetry

func Init() {
	config.Init()
	c := config.GetConfig()
	_global = new(c)
}

func Get() *pkg_telemetry.Telemetry {
	if _global == nil {
		log.Fatalf("telemetry not initialized")
	}
	return _global
}

func new(config config.Config) *pkg_telemetry.Telemetry {
	traceProvider := config.Telemetry.TraceProvider
	traceURL := config.Telemetry.TraceURL
	telemetryConfig, err := pkg_telemetry.NewTelemetryConfig(traceProvider, traceURL)
	if err != nil {
		log.Fatalf("error creating telemetry config: %v", err)
	}

	t, err := pkg_telemetry.New(telemetryConfig)
	if err != nil {
		log.Fatalf("error creating telemetry: %v", err)
	}

	return t
}
