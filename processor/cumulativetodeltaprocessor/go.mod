module github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.109.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.109.1-0.20240923174701-20f73e2294b5
	go.opentelemetry.io/collector/confmap v1.15.1-0.20240923174701-20f73e2294b5
	go.opentelemetry.io/collector/consumer v0.109.1-0.20240923174701-20f73e2294b5
	go.opentelemetry.io/collector/consumer/consumertest v0.109.1-0.20240923174701-20f73e2294b5
	go.opentelemetry.io/collector/pdata v1.15.1-0.20240923174701-20f73e2294b5
	go.opentelemetry.io/collector/processor v0.109.1-0.20240923174701-20f73e2294b5
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.109.1-0.20240923174701-20f73e2294b5 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.109.1-0.20240923174701-20f73e2294b5 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.109.1-0.20240923174701-20f73e2294b5 // indirect
	go.opentelemetry.io/collector/internal/globalsignal v0.0.0-20240923174701-20f73e2294b5 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.109.1-0.20240923174701-20f73e2294b5 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.109.1-0.20240923174701-20f73e2294b5 // indirect
	go.opentelemetry.io/collector/pipeline v0.0.0-20240923174701-20f73e2294b5 // indirect
	go.opentelemetry.io/collector/processor/processorprofiles v0.109.1-0.20240923174701-20f73e2294b5 // indirect
	go.opentelemetry.io/otel v1.30.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/sdk v1.30.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240822170219-fc7c04adadcd // indirect
	google.golang.org/grpc v1.66.2 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter => ../../internal/filter

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl => ../../pkg/ottl

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
