module github.com/open-telemetry/opentelemetry-collector-contrib/extension/httpforwarder

go 1.16

require (
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.18.1
)

replace go.opentelemetry.io/collector => go.opentelemetry.io/collector v0.29.1-0.20210716020257-4d8e3082465d

replace go.opentelemetry.io/collector/model => go.opentelemetry.io/collector/model v0.0.0-20210716020257-4d8e3082465d
