package config

type Telemetry struct {
	Enabled       bool   `env:"TRACE_ENABLED,required"`
	TraceProvider string `env:"TRACE_PROVIDER,required"`
	TraceURL      string `env:"TRACE_URL,required"`
}
