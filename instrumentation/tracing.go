package instrumentation

import (
	"context"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

func InitTracer() (func(context.Context) error, error) {
	ctx := context.Background()

	// Set up OTLP exporter
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if endpoint == "" {
		endpoint = "otel-collector:4317"
	}

	conn, err := grpc.DialContext(ctx, endpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	// Create trace provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("phonebook-api"),
		)),
	)

	// Set global tracer provider
	otel.SetTracerProvider(tp)

	return func(ctx context.Context) error {
		if err := tp.Shutdown(ctx); err != nil {
			return err
		}
		if err := conn.Close(); err != nil {
			return err
		}
		return nil
	}, nil
}
