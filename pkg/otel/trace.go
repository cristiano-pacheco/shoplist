package telemetry

import (
	"fmt"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func newTracerProvider(config TelemetryConfig, res *resource.Resource) (*sdktrace.TracerProvider, error) {
	if !config.TraceEnabled {
		return sdktrace.NewTracerProvider(), nil
	}

	exp, err := newExporter(config)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
	)

	return tp, nil
}

func newExporter(config TelemetryConfig) (sdktrace.SpanExporter, error) {
	// Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
	if config.TracerVendor.String() == TraceVendorZipkin {
		return zipkin.New(
			config.TraceURL,
		)
	}

	if config.TracerVendor.String() == TraceVendorStdout {
		return stdouttrace.New(
			stdouttrace.WithPrettyPrint(),
		)
	}

	return nil, fmt.Errorf("invalid trace provider: %s", config.TracerVendor.String())
}
