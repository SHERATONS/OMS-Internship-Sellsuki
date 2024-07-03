package Observability

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"

	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"log"
)

func InitTracer() func() {
	ctx := context.Background()
	traceExporter, err := otlptrace.New(ctx, otlptracehttp.NewClient(
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint("localhost:4318"),
	))
	if err != nil {
		log.Fatal(err)
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("OMS"))

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res))

	otel.SetTracerProvider(tp)

	return nil
}
