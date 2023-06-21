package tracing

import (
	"context"
	"os"
	"time"

	"github.com/erfanmomeniii/user-management/internal/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.16.0"
	"go.opentelemetry.io/otel/trace"
)

func Init(exporter func(*config.Config) (traceSdk.SpanExporter, error), cfg *config.Config) (*traceSdk.TracerProvider, trace.Tracer, error) {
	if !cfg.Tracer.Enabled {
		return nil, trace.NewNoopTracerProvider().Tracer(config.AppName), nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, nil, err
	}

	exp, err := exporter(cfg)
	if err != nil {
		return nil, nil, err
	}

	tracerProvider := traceSdk.NewTracerProvider(
		traceSdk.WithBatcher(exp),
		traceSdk.WithSampler(traceSdk.TraceIDRatioBased(cfg.Tracer.SamplerRatio)),
		traceSdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceInstanceIDKey.String(hostname),
			semconv.ServiceNameKey.String(config.AppName),
		)),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	return tracerProvider, otel.Tracer(config.AppName), nil
}

func Close(provider *traceSdk.TracerProvider) error {
	if provider != nil {
		c, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := provider.Shutdown(c); err != nil {
			return err
		}
	}

	return nil
}
