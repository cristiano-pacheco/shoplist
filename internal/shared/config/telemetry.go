package config

type Telemetry struct {
	TraceProvider string `env:"TRACE_PROVIDER,required"`
	TraceURL      string `env:"TRACE_URL,required"`
}
