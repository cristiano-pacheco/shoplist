package telemetry

type TelemetryConfig struct {
	TraceEnabled  bool
	TraceProvider TraceProvider
	TraceURL      string
}

func NewTelemetryConfig(traceEnabled bool, traceProvider, traceURL string) (TelemetryConfig, error) {
	tp, err := NewTraceProvider(traceProvider)
	if err != nil {
		return TelemetryConfig{}, err
	}

	tc := TelemetryConfig{
		TraceEnabled:  traceEnabled,
		TraceProvider: tp,
		TraceURL:      traceURL,
	}

	return tc, nil
}
