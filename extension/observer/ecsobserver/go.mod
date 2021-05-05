module github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver

go 1.15

require (
	github.com/aws/aws-sdk-go v1.38.30
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.25.1-0.20210504213219-970b76cc794a
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
)

// WIP update for otelcol changes
replace go.opentelemetry.io/collector => github.com/pmatyjasek-sumo/opentelemetry-collector v0.25.1-0.20210428081312-72ef9d6ccfe5
