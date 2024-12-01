package telemetry

import (
	"log"

	"github.com/cristiano-pacheco/shoplist/internal/shared/config"
	pkg_telemetry "github.com/cristiano-pacheco/shoplist/pkg/telemetry"
)

var _global *pkg_telemetry.Telemetry

func Init(cfg config.Config) {
	_global = new(cfg)
}

func Get() *pkg_telemetry.Telemetry {
	if _global == nil {
		log.Fatalf("telemetry not initialized")
	}
	return _global
}

func new(config config.Config) *pkg_telemetry.Telemetry {
	if config.Telemetry.Enabled && config.Telemetry.TracerVendor == "" {
		log.Fatalf("setting up telemetry: TELEMETRY_TRACER_VENDOR is not set")
		return nil
	}

	if config.Telemetry.Enabled && config.Telemetry.TracerURL == "" {
		log.Fatalf("setting up telemetry: TELEMETRY_TRACER_URL is not set")
		return nil
	}

	tracerVendor, err := pkg_telemetry.NewTracerVendor(config.Telemetry.TracerVendor)
	if err != nil {
		log.Fatalf("error creating tracer vendor: %v", err)
	}

	tc := pkg_telemetry.TelemetryConfig{
		AppName:      config.App.Name,
		AppVersion:   config.App.Version,
		TraceEnabled: config.Telemetry.Enabled,
		TracerVendor: tracerVendor,
		TraceURL:     config.Telemetry.TracerURL,
	}

	return pkg_telemetry.New(tc)
}
