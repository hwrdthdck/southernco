module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter

go 1.15

require (
	github.com/DataDog/datadog-agent v0.0.0-20200417180928-f454c60bc16f
	github.com/DataDog/viper v1.8.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/klauspost/compress v1.10.10
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/stretchr/testify v1.6.1
	github.com/zorkian/go-datadog-api v2.29.0+incompatible // indirect
	go.opencensus.io v0.22.4
	go.opentelemetry.io/collector v0.12.1-0.20201012183541-526f34200197
	go.uber.org/zap v1.16.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.26.0
	gopkg.in/zorkian/go-datadog-api.v2 v2.29.0
)
