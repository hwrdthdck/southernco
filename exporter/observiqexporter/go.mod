module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/observiqexporter

go 1.16

require (
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.29.1-0.20210630154722-7d0a0398174e
	go.uber.org/zap v1.18.1
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

replace go.opentelemetry.io/collector/model => go.opentelemetry.io/collector/model v0.0.0-20210630154722-7d0a0398174e
