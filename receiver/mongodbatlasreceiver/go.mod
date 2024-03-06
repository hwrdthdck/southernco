module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver

go 1.21

require (
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/google/go-cmp v0.6.0
	github.com/mongodb-forks/digest v1.0.5
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.96.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.96.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.96.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.96.0
	github.com/stretchr/testify v1.9.0
	go.mongodb.org/atlas v0.36.0
	go.opentelemetry.io/collector/component v0.96.1-0.20240306115632-b2693620eff6
	go.opentelemetry.io/collector/config/configopaque v1.3.1-0.20240306115632-b2693620eff6
	go.opentelemetry.io/collector/config/configretry v0.96.1-0.20240306115632-b2693620eff6
	go.opentelemetry.io/collector/config/configtls v0.96.1-0.20240306115632-b2693620eff6
	go.opentelemetry.io/collector/confmap v0.96.1-0.20240306115632-b2693620eff6
	go.opentelemetry.io/collector/consumer v0.96.1-0.20240306115632-b2693620eff6
	go.opentelemetry.io/collector/extension v0.96.1-0.20240306115632-b2693620eff6
	go.opentelemetry.io/collector/pdata v1.3.1-0.20240306115632-b2693620eff6
	go.opentelemetry.io/collector/receiver v0.96.1-0.20240306115632-b2693620eff6
	go.opentelemetry.io/otel/metric v1.24.0
	go.opentelemetry.io/otel/trace v1.24.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/expr-lang/expr v1.16.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/influxdata/go-syslog/v3 v3.0.1-0.20230911200830-875f5bc594a4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.0 // indirect
	github.com/leodido/ragel-machinery v0.0.0-20181214104525-299bdde78165 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.96.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.96.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	go.opentelemetry.io/collector v0.96.1-0.20240306115632-b2693620eff6 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.96.1-0.20240306115632-b2693620eff6 // indirect
	go.opentelemetry.io/collector/featuregate v1.3.1-0.20240306115632-b2693620eff6 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.46.0 // indirect
	go.opentelemetry.io/otel/sdk v1.24.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.24.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gonum.org/v1/gonum v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/grpc v1.62.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza => ../../pkg/stanza

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage => ../../extension/storage

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
