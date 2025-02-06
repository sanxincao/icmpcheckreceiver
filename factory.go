package icmpcheckreceiver

import (
	"context"
	"errors"

	"github.com/sanxincao/icmpcheckreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability),
	)
}

func createDefaultConfig() component.Config {
	cfg := scraperhelper.NewDefaultControllerConfig()

	return &Config{
		ControllerConfig: cfg,
		Targets:          []Target{},
	}
}

func createMetricsReceiver(
	ctx context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	receiverCfg, ok := cfg.(*Config)
	if !ok {
		return nil, errors.New("config is not a valid icmpping receiver config")
	}

	icmpScraper, err := newScraper(set.Logger, receiverCfg.Targets)
	if err != nil {
		return nil, err
	}

	scraper, err := scraper.NewMetrics(icmpScraper.Scrape, scraper.WithStart(icmpScraper.ping))
	if err != nil {
		return nil, err
	}

	return scraperhelper.NewMetricsController(&receiverCfg.ControllerConfig, set, nextConsumer, scraperhelper.AddScraper(metadata.Type, scraper))

}
