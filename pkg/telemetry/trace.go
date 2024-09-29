package telemetry

import (
	"fmt"
	"os"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func newTraceProvider(config TelemetryConfig) (*sdktrace.TracerProvider, error) {
	if !config.TraceEnabled {
		return sdktrace.NewTracerProvider(), nil
	}

	exp, err := newExporter(config)
	if err != nil {
		return nil, err
	}

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "gomodulith"
	}

	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(appName),
		),
	)

	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

	return tp, nil
}

func newExporter(config TelemetryConfig) (sdktrace.SpanExporter, error) {
	// Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
	if config.TraceProvider.String() == TraceProviderZipkin {
		return zipkin.New(
			config.TraceURL,
		)
	}

	if config.TraceProvider.String() == TraceProviderStdout {
		return stdouttrace.New(
			stdouttrace.WithPrettyPrint(),
		)
	}

	return nil, fmt.Errorf("invalid trace provider: %s", config.TraceProvider.String())
}
