module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/jaegertracing/jaeger v1.22.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.0.0-00010101000000-000000000000
	github.com/signalfx/sapm-proto v0.7.0
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.23.0
	go.opentelemetry.io/collector v0.26.1-0.20210510162429-51281a719256
	go.uber.org/zap v1.16.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => ../../internal/splunk
