package utils

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"google.golang.org/grpc/encoding/gzip"
)

func InitTracer(ctx context.Context, serviceName string, serviceVersion string, jaegerURL string) *otlptrace.Exporter {
	resources, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		),
		resource.WithProcessRuntimeName(),
		resource.WithProcessRuntimeVersion(),
	)
	if err != nil {
		log.Fatalf("Failed creating resources for tracer : %v", err)
	}

	traceExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithCompressor(gzip.Name),
		otlptracegrpc.WithEndpoint(jaegerURL),
	)
	if err != nil {
		log.Fatalf("Failed creating trace exporter : %v", err)
	}

	otel.SetTracerProvider(trace.NewTracerProvider(
		trace.WithSpanProcessor(trace.NewBatchSpanProcessor(traceExporter)),
		trace.WithResource(resources),
	))

	return traceExporter
}
