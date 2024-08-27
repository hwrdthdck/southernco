module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver

go 1.22.0

require (
	github.com/gosnmp/gosnmp v1.38.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.107.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.107.0
	github.com/stretchr/testify v1.9.0
	github.com/testcontainers/testcontainers-go v0.33.0
	go.opentelemetry.io/collector/component v0.107.1-0.20240827012220-5963d446ca4a
	go.opentelemetry.io/collector/config/configopaque v1.13.1-0.20240827012220-5963d446ca4a
	go.opentelemetry.io/collector/confmap v0.107.1-0.20240827012220-5963d446ca4a
	go.opentelemetry.io/collector/consumer v0.107.1-0.20240827012220-5963d446ca4a
	go.opentelemetry.io/collector/consumer/consumertest v0.107.1-0.20240827012220-5963d446ca4a
	go.opentelemetry.io/collector/otelcol/otelcoltest v0.107.1-0.20240827012220-5963d446ca4a
	go.opentelemetry.io/collector/pdata v1.13.1-0.20240827012220-5963d446ca4a
	go.opentelemetry.io/collector/receiver v0.107.1-0.20240827012220-5963d446ca4a
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/containerd/containerd v1.7.18 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/platforms v0.2.1 // indirect
	github.com/cpuguy83/dockercfg v0.3.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/docker v27.1.1+incompatible // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-viper/mapstructure/v2 v2.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/sys/user v0.1.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.107.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/client_golang v1.20.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/shirou/gopsutil/v4 v4.24.7 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/collector v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/component/componentprofiles v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/confmap/provider/envprovider v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/confmap/provider/fileprovider v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/confmap/provider/httpprovider v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/confmap/provider/yamlprovider v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/connector v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/exporter v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/extension v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/featuregate v1.13.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/internal/globalgates v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/otelcol v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/processor v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/semconv v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/collector/service v0.107.1-0.20240827012220-5963d446ca4a // indirect
	go.opentelemetry.io/contrib/config v0.8.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.53.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.28.0 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.4.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.50.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.28.0 // indirect
	go.opentelemetry.io/otel/log v0.4.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.4.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	gonum.org/v1/gonum v0.15.1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
