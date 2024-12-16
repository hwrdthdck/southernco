module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter

go 1.22.0

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/fsnotify/fsnotify v1.8.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/snappy v0.0.4
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.115.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.115.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.115.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus v0.115.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite v0.115.0
	github.com/prometheus/prometheus v0.54.1
	github.com/stretchr/testify v1.10.0
	github.com/tidwall/wal v1.1.8
	go.opentelemetry.io/collector/component v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/component/componenttest v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/config/confighttp v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/config/configopaque v1.21.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/config/configretry v1.21.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/config/configtelemetry v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/config/configtls v1.21.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/confmap v1.21.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/consumer/consumererror v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/exporter v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/exporter/exportertest v0.115.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/featuregate v1.21.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/collector/pdata v1.21.1-0.20241216091623-8ac40a01a5ff
	go.opentelemetry.io/otel v1.32.0
	go.opentelemetry.io/otel/metric v1.32.0
	go.opentelemetry.io/otel/sdk/metric v1.32.0
	go.opentelemetry.io/otel/trace v1.32.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.13.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.2 // indirect
	github.com/alecthomas/units v0.0.0-20240626203959-61d1e3462e30 // indirect
	github.com/aws/aws-sdk-go v1.54.19 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/regexp v0.0.0-20240518133315-a468a5bfb3bc // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.20.4 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.61.0 // indirect
	github.com/prometheus/common/sigv4 v0.1.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/tidwall/gjson v1.10.2 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/tinylru v1.1.0 // indirect
	go.opentelemetry.io/collector/client v1.21.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/config/configauth v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/config/configcompression v1.21.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/config/internal v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/consumer v1.21.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.0.0-20241215143820-6147243aaaa1 // indirect
	go.opentelemetry.io/collector/exporter/xexporter v0.0.0-20241215143820-6147243aaaa1 // indirect
	go.opentelemetry.io/collector/extension v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/extension/auth v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/pipeline v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/receiver v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.0.0-20241215143820-6147243aaaa1 // indirect
	go.opentelemetry.io/collector/semconv v0.115.1-0.20241216091623-8ac40a01a5ff // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/oauth2 v0.24.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241206012308-a4fef0638583 // indirect
	google.golang.org/grpc v1.68.1 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apimachinery v0.29.3 // indirect
	k8s.io/client-go v0.29.3 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry => ../../pkg/resourcetotelemetry

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus => ../../pkg/translator/prometheus

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite => ../../pkg/translator/prometheusremotewrite

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
