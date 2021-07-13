module github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver

go 1.16

require (
	github.com/aws/aws-sdk-go v1.39.5
	github.com/hashicorp/golang-lru v0.5.4
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.29.1-0.20210712235908-f9dacb8402fe
	go.uber.org/multierr v1.7.0
	go.uber.org/zap v1.18.1
	gopkg.in/yaml.v2 v2.4.0
)

replace go.opentelemetry.io/collector/model => go.opentelemetry.io/collector/model v0.0.0-20210712235908-f9dacb8402fe
