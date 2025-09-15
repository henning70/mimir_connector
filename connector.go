package validationconnector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

var (
	Type = component.MustNewType("validationconnector")
)

const (
	TracesToTracesStability   = component.StabilityLevelBeta
	MetricsToMetricsStability = component.StabilityLevelBeta
	LogsToLogsStability       = component.StabilityLevelBeta
)

type Config struct{}

type validate struct {
	consumer.Traces
	consumer.Metrics
	consumer.Logs
	component.StartFunc
	component.ShutdownFunc
}

// NewFactory creates a factory for example connector.
func NewFactory() connector.Factory {
	// OpenTelemetry connector factory to make a factory for connectors
	return connector.NewFactory(
		Type,
		createDefaultConfig,
		connector.WithTracesToTraces(createTracesToTraces, TracesToTracesStability),
		connector.WithMetricsToMetrics(createMetricsToMetrics, MetricsToMetricsStability),
		connector.WithLogsToLogs(createLogsToLogs, LogsToLogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

// createTracesToTraces creates a trace receiver based on provided config.
func createTracesToTraces(
	_ context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer consumer.Traces,
) (connector.Traces, error) {
	return &validate{Traces: nextConsumer}, nil
}

// createMetricsToMetrics creates a metrics receiver based on provided config.
func createMetricsToMetrics(
	context context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer consumer.Metrics,
) (connector.Metrics, error) {
	printIt(context)
	return &validate{Metrics: nextConsumer}, nil
}

// createLogsToLogs creates a log receiver based on provided config.
func createLogsToLogs(
	_ context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer consumer.Logs,
) (connector.Logs, error) {
	return &validate{Logs: nextConsumer}, nil
}

func (c *validate) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func printIt(in context.Context) {
	fmt.Println(in)
}
