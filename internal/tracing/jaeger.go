package tracing

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/erfanmomeniii/user-management/internal/config"
)

func InitJaeger(cfg *config.Config) (traceSdk.SpanExporter, error) {
	return jaeger.New(jaeger.WithAgentEndpoint(
		jaeger.WithAgentHost(cfg.Tracing.AgentHost),
		jaeger.WithAgentPort(cfg.Tracing.AgentPort),
	))
}
