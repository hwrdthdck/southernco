module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.118.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/consumer v1.24.1-0.20250203014413-643a35ffbcea
	go.opentelemetry.io/collector/exporter v0.118.1-0.20250131104636-a737a48402e0
	go.opentelemetry.io/collector/pdata v1.24.1-0.20250203014413-643a35ffbcea
	go.uber.org/goleak v1.3.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/collector/component v0.118.1-0.20250131104636-a737a48402e0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.118.1-0.20250131104636-a737a48402e0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.118.1-0.20250131104636-a737a48402e0 // indirect
	go.opentelemetry.io/collector/pipeline v0.118.1-0.20250131104636-a737a48402e0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/grpc v1.70.0 // indirect
	google.golang.org/protobuf v1.36.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../golden
