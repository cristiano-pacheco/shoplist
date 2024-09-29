package telemetry

import (
	"context"
	"os"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel/trace"
)

type Telemetry struct {
	tracerProvider trace.TracerProvider
}

func New(config TelemetryConfig) (*Telemetry, error) {
	traceProvider, err := newTraceProvider(config)
	if err != nil {
		return nil, err
	}

	t := Telemetry{
		tracerProvider: traceProvider,
	}

	return &t, nil
}

func (t *Telemetry) StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "gomodulith"
	}
	tracer := t.tracerProvider.Tracer(appName)
	return tracer.Start(ctx, name)
}

func (t *Telemetry) End(ctx context.Context) {
	if tp, ok := t.tracerProvider.(*sdktrace.TracerProvider); ok {
		tp.Shutdown(ctx)
	}
}
