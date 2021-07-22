module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter

go 1.16

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.30.2-0.20210721214806-6cce4224c6e7
	go.opentelemetry.io/collector/model v0.30.2-0.20210721214806-6cce4224c6e7
	go.uber.org/zap v1.18.1
	google.golang.org/protobuf v1.27.1

)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => ../../internal/splunk
