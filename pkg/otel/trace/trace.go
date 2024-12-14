package trace

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type Trace interface {
	StartSpan(ctx context.Context, name string) (context.Context, oteltrace.Span)
	Shutdown(ctx context.Context) error
}

type trace struct {
	tracer         oteltrace.Tracer
	tracerProvider *sdktrace.TracerProvider
}

type TracerConfig struct {
	AppName      string
	AppVersion   string
	TracerVendor string
	TraceURL     string
	TraceEnabled bool
}

func New(config TracerConfig) Trace {
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
	tp := newTracerProvider(config, res)
	otel.SetTracerProvider(tp)
	// Set up propagator.
	otel.SetTextMapPropagator(newPropagator())
	t := tp.Tracer(config.AppName)
	trace := trace{
		tracer: t,
	}
	return &trace
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTracerProvider(config TracerConfig, res *resource.Resource) *sdktrace.TracerProvider {
	if !config.TraceEnabled {
		return sdktrace.NewTracerProvider()
	}
	exp := newExporter(config)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
	)
	return tp
}

func newExporter(config TracerConfig) sdktrace.SpanExporter {
	exporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(config.TraceURL),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	return exporter
}

func (t *trace) StartSpan(ctx context.Context, name string) (context.Context, oteltrace.Span) {
	return t.tracer.Start(ctx, name)
}

func (t *trace) Shutdown(ctx context.Context) error {
	return t.tracerProvider.Shutdown(ctx)
}
