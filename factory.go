package icmpcheckreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

var (
	typeStr = component.MustNewType("icmpcheckreceiver")
)

const (
	defaultInterval = 1 * time.Minute
)

func createDefaultConfig() component.Config {
	return &Config{
		Interval: defaultInterval,
	}
}

func createMetricsReceiver(_ context.Context, params receiver.Settings, baseCfg component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	logger := params.Logger
	cfg := baseCfg.(*Config)

	icmpRcvr := newIcmpCheckReceiver(logger, cfg, consumer)

	return receiverhelper.NewReceiver(&cfg.ReceiverSettings, params, consumer, receiverhelper.AddScraper(icmpRcvr))
}

// NewFactory creates a factory for icmpcheckreceiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, component.StabilityLevelAlpha))
}
