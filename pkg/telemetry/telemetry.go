package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type TememetryI interface {
	StartSpan(ctx context.Context, name string) (context.Context, trace.Span)
	Shutdown(ctx context.Context) error
}

type Telemetry struct {
	tracer         trace.Tracer
	tracerProvider *sdktrace.TracerProvider
}

func New(config TelemetryConfig) *Telemetry {
	// Set up resource.
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.AppName),
		),
	)

	if err != nil {
		panic(err)
	}

	// Set up tracer provider.
	tp, err := newTracerProvider(config, res)
	if err != nil {
		panic(err)
	}

	otel.SetTracerProvider(tp)

	// Set up propagator.
	otel.SetTextMapPropagator(newPropagator())

	t := tp.Tracer("github.com/cristiano-pacheco/gomodulith")

	telemetry := Telemetry{
		tracer: t,
	}

	return &telemetry
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func (t *Telemetry) StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, name)
}

func (t *Telemetry) Shutdown(ctx context.Context) error {
	return t.tracerProvider.Shutdown(ctx)
}
