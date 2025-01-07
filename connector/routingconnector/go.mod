module github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.116.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.116.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/client v1.22.1-0.20250107062214-ced38e8af2ae
	go.opentelemetry.io/collector/component v0.116.1-0.20250107062214-ced38e8af2ae
	go.opentelemetry.io/collector/component/componenttest v0.116.1-0.20250107062214-ced38e8af2ae
	go.opentelemetry.io/collector/confmap v1.22.1-0.20250107062214-ced38e8af2ae
	go.opentelemetry.io/collector/connector v0.116.1-0.20250107062214-ced38e8af2ae
	go.opentelemetry.io/collector/connector/connectortest v0.116.1-0.20250107062214-ced38e8af2ae
	go.opentelemetry.io/collector/consumer v1.22.1-0.20250107062214-ced38e8af2ae
	go.opentelemetry.io/collector/consumer/consumertest v0.116.1-0.20250107062214-ced38e8af2ae
	go.opentelemetry.io/collector/pdata v1.22.1-0.20250107062214-ced38e8af2ae
	go.opentelemetry.io/collector/pipeline v0.116.1-0.20250107062214-ced38e8af2ae
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.69.2
)

require (
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/antchfx/xmlquery v1.4.3 // indirect
	github.com/antchfx/xpath v1.3.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/elastic/go-grok v0.3.1 // indirect
	github.com/elastic/lunes v0.1.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
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
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.116.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.116.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/ua-parser/uap-go v0.0.0-20240611065828-3a4781585db6 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.116.1-0.20250107062214-ced38e8af2ae // indirect
	go.opentelemetry.io/collector/connector/xconnector v0.116.1-0.20250107062214-ced38e8af2ae // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.116.1-0.20250107062214-ced38e8af2ae // indirect
	go.opentelemetry.io/collector/internal/fanoutconsumer v0.116.1-0.20250107062214-ced38e8af2ae // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.116.1-0.20250107062214-ced38e8af2ae // indirect
	go.opentelemetry.io/collector/pipeline/xpipeline v0.116.1-0.20250107062214-ced38e8af2ae // indirect
	go.opentelemetry.io/collector/semconv v0.116.1-0.20250107062214-ced38e8af2ae // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl => ../../pkg/ottl

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden

replace go.opentelemetry.io/collector/scraper/scraperhelper v0.116.0 => go.opentelemetry.io/collector/scraper/scraperhelper v0.0.0-20250107062214-ced38e8af2ae

replace go.opentelemetry.io/collector/extension/xextension v0.116.0 => go.opentelemetry.io/collector/extension/xextension v0.0.0-20250107062214-ced38e8af2ae
