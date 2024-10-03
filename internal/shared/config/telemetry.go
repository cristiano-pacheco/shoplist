package config

type Telemetry struct {
	Enabled      bool   `env:"TRACER_ENABLED,required"`
	TracerVendor string `env:"TRACER_VENDOR"`
	TracerURL    string `env:"TRACER_URL"`
}
