module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter

go 1.16

require (
	contrib.go.opencensus.io/exporter/stackdriver v0.13.8
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v1.0.0-RC2
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.23.0
	go.opentelemetry.io/collector v0.33.1-0.20210827152330-09258f969908
	go.opentelemetry.io/collector/model v0.33.1-0.20210827152330-09258f969908
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
	go.opentelemetry.io/otel/trace v1.0.0-RC2
	google.golang.org/api v0.54.0
	google.golang.org/genproto v0.0.0-20210805201207-89edb61ffb67
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus => ../../pkg/translator/opencensus

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal
