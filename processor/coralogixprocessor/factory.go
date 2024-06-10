package coralogixprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/coralogixprocessor"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/coralogixprocessor/internal/metadata"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

// NewFactory returns a new factory for the Span processor.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		metadata.Type,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, component.StabilityLevelDevelopment))
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createTracesProcessor(
	ctx context.Context,
	params processor.CreateSettings,
	baseCfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	coralogixsCfg := baseCfg.(*Config)

	coralogixProcessor, err := newCoralogixProcessor(ctx,
		params,
		coralogixsCfg,
		nextConsumer)
	if err != nil {
		return nil, err
	}

	return coralogixProcessor, nil
}
