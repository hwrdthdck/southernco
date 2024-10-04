//module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure

go 1.22.0

require (
	github.com/json-iterator/go v1.1.12
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.110.0
	github.com/relvacode/iso8601 v1.4.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.110.1-0.20241004063257-d6cd5935eefc
	go.opentelemetry.io/collector/pdata v1.16.1-0.20241004063257-d6cd5935eefc
	go.opentelemetry.io/collector/semconv v0.110.1-0.20241004063257-d6cd5935eefc
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.110.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.110.1-0.20241004063257-d6cd5935eefc // indirect
	go.opentelemetry.io/collector/internal/globalsignal v0.110.1-0.20241004063257-d6cd5935eefc // indirect
	go.opentelemetry.io/collector/pipeline v0.110.1-0.20241004063257-d6cd5935eefc // indirect
	go.opentelemetry.io/otel v1.30.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/grpc v1.67.1 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pdatatest

module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../golden
