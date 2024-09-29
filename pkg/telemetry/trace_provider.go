package telemetry

import "fmt"

type TraceProvider struct {
	value string
}

const (
	TraceProviderZipkin = "zipkin"
	TraceProviderStdout = "stdout"
)

func NewTraceProvider(s string) (TraceProvider, error) {
	err := validateTraceProvider(s)
	if err != nil {
		return TraceProvider{}, err
	}

	return TraceProvider{value: s}, nil
}

func validateTraceProvider(s string) error {
	switch s {
	case TraceProviderZipkin, TraceProviderStdout:
		return nil
	default:
		return fmt.Errorf("invalid trace provider: %s", s)
	}
}

func (t TraceProvider) String() string {
	return t.value
}
