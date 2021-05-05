module github.com/open-telemetry/opentelemetry-collector-contrib/testbed/mockdatareceivers/mockawsxrayreceiver

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/pelletier/go-toml v1.8.0 // indirect
	go.opentelemetry.io/collector v0.25.1-0.20210504213736-b8824d020718
	go.uber.org/zap v1.16.0
	gopkg.in/ini.v1 v1.57.0 // indirect
)
// WIP update for otelcol changes
replace go.opentelemetry.io/collector => github.com/pmatyjasek-sumo/opentelemetry-collector v0.26.1-0.20210505092123-44f32bb740c4
