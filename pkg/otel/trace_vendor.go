package telemetry

import (
	"fmt"
)

type TracerVendor struct {
	value string
}

const (
	TraceVendorZipkin = "zipkin"
	TraceVendorStdout = "stdout"
)

func NewTracerVendor(s string) (TracerVendor, error) {
	err := validateTracerVendor(s)
	if err != nil {
		return TracerVendor{}, err
	}

	return TracerVendor{value: s}, nil
}

func validateTracerVendor(s string) error {
	switch s {
	case TraceVendorZipkin, TraceVendorStdout:
		return nil
	default:
		return fmt.Errorf("invalid trace provider: %s", s)
	}
}

func (t TracerVendor) String() string {
	return t.value
}
