module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zookeeperreceiver

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.15.1-0.20201125171618-60498105d42f
	go.uber.org/zap v1.16.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common
