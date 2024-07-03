package googlecloudmonitoringreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudmonitoringreceiver"

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type monitoringReceiver struct {
	config *Config
	logger *zap.Logger
	cancel context.CancelFunc
}

func newGoogleCloudMonitoringReceiver(cfg *Config, logger *zap.Logger) *monitoringReceiver {
	return &monitoringReceiver{
		config: cfg,
		logger: logger,
	}
}

func (m *monitoringReceiver) Scrape(ctx context.Context) (pmetric.Metrics, error) {
	metrics := pmetric.NewMetrics()
	m.logger.Debug("Scrape metrics ")

	return metrics, nil
}

func (m *monitoringReceiver) Start(ctx context.Context, _ component.Host) error {
	ctx, m.cancel = context.WithCancel(ctx)
	err := m.initialize(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *monitoringReceiver) Shutdown(context.Context) error {
	m.cancel()
	m.logger.Debug("shutting down googlecloudmonitoringreceiver receiver")

	return nil
}

func (m *monitoringReceiver) initialize(ctx context.Context) error {
	// Implement the logic for handling metrics here.
	return nil
}
