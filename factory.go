package icmpcheckreceiver

import (
	"context"

	"github.com/sanxincao/icmpcheckreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Interval: defaultInterval,
	}
}

func createMetricsReceiver(_ context.Context, params receiver.Settings, baseCfg component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	logger := params.Logger
	cfg := baseCfg.(*Config)

	icmpRcvr := newIcmpCheckReceiver(logger, cfg, consumer)

	return icmpRcvr, nil
}

// NewFactory creates a factory for icmpcheckreceiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, component.StabilityLevelAlpha))
}
