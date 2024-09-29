package telemetry

type TelemetryConfig struct {
	TraceProvider TraceProvider
	TraceURL      string
}

func NewTelemetryConfig(traceProvider, traceURL string) (TelemetryConfig, error) {
	tp, err := NewTraceProvider(traceProvider)
	if err != nil {
		return TelemetryConfig{}, err
	}
	return TelemetryConfig{
		TraceProvider: tp,
		TraceURL:      traceURL,
	}, nil
}
