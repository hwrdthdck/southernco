module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter

go 1.22.0

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/gobwas/glob v0.2.3
	github.com/gogo/protobuf v1.3.2
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.113.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.113.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.113.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.113.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.113.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.113.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.113.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx v0.113.0
	github.com/shirou/gopsutil/v4 v4.24.10
	github.com/signalfx/com_signalfx_metrics_protobuf v0.0.3
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.113.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/config/confighttp v0.113.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/config/configopaque v1.19.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/config/configretry v1.19.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/config/configtls v1.19.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/confmap v1.19.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/consumer v0.113.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/consumer/consumererror v0.113.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/exporter v0.113.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/exporter/exportertest v0.113.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/pdata v1.19.1-0.20241115165626-8b99b8023ca3
	go.opentelemetry.io/collector/semconv v0.113.1-0.20241115165626-8b99b8023ca3
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/sys v0.26.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ebitengine/purego v0.8.1 // indirect
	github.com/elastic/lunes v0.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.113.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/collector/client v1.19.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/configauth v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.19.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/config/internal v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/exporter/exporterprofiles v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/extension v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/extension/auth v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/pipeline v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/receiver v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/receiver/receiverprofiles v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.113.1-0.20241115165626-8b99b8023ca3 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.31.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240822170219-fc7c04adadcd // indirect
	google.golang.org/grpc v1.67.1 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => ../../internal/splunk

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => ../../pkg/batchperresourceattr

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata => ../../pkg/experimentalmetricmetadata

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx => ../../pkg/translator/signalfx

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

// https://github.com/go-openapi/spec/issues/156
replace github.com/go-openapi/spec v0.20.5 => github.com/go-openapi/spec v0.20.6

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
