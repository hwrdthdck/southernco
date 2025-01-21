module github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.117.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter v0.117.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.117.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.117.1-0.20250117002813-e970f8bb1258
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.118.1-0.20250121144026-bc76c3284db9
	go.opentelemetry.io/collector/component/componenttest v0.118.1-0.20250121144026-bc76c3284db9
	go.opentelemetry.io/collector/confmap v1.24.1-0.20250121144026-bc76c3284db9
	go.opentelemetry.io/collector/consumer v1.24.1-0.20250121144026-bc76c3284db9
	go.opentelemetry.io/collector/consumer/consumertest v0.118.1-0.20250121144026-bc76c3284db9
	go.opentelemetry.io/collector/pdata v1.24.1-0.20250121144026-bc76c3284db9
	go.opentelemetry.io/collector/processor v0.118.1-0.20250121144026-bc76c3284db9
	go.opentelemetry.io/collector/processor/processortest v0.118.1-0.20250121144026-bc76c3284db9
	go.opentelemetry.io/collector/semconv v0.118.1-0.20250121144026-bc76c3284db9
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/antchfx/xmlquery v1.4.3 // indirect
	github.com/antchfx/xpath v1.3.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elastic/go-grok v0.3.1 // indirect
	github.com/elastic/lunes v0.1.0 // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.117.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/ua-parser/uap-go v0.0.0-20240611065828-3a4781585db6 // indirect
	go.opentelemetry.io/collector/client v1.24.1-0.20250121144026-bc76c3284db9 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.118.1-0.20250121144026-bc76c3284db9 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.118.1-0.20250121144026-bc76c3284db9 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.118.1-0.20250121144026-bc76c3284db9 // indirect
	go.opentelemetry.io/collector/featuregate v1.24.1-0.20250121144026-bc76c3284db9 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.118.1-0.20250121144026-bc76c3284db9 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.118.1-0.20250121144026-bc76c3284db9 // indirect
	go.opentelemetry.io/collector/pipeline v0.118.1-0.20250121144026-bc76c3284db9 // indirect
	go.opentelemetry.io/collector/processor/xprocessor v0.118.1-0.20250121144026-bc76c3284db9 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	google.golang.org/grpc v1.69.4 // indirect
	google.golang.org/protobuf v1.36.3 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter => ../../internal/filter

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl => ../../pkg/ottl

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
