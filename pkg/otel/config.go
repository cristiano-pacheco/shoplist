package telemetry

type TelemetryConfig struct {
	AppName      string
	AppVersion   string
	TraceEnabled bool
	TracerVendor TracerVendor
	TraceURL     string
}
