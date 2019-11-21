module github.com/open-telemetry/opentelemetry-collector-contrib

go 1.12

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter v0.0.0 => ./exporter/azuremonitorexporter

require (
	github.com/client9/misspell v0.3.4
	github.com/google/addlicense v0.0.0-20190907113143-be125746c2c4
	github.com/open-telemetry/opentelemetry-collector v0.2.1-0.20191126183205-e94dd19191e0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter v0.0.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stackdriverexporter v0.0.0-20191126142441-b2a048090ad6
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinscribereceiver v0.0.0-20191126142441-b2a048090ad6
	github.com/pavius/impi v0.0.0-20180302134524-c1cbdcb8df2b
	golang.org/x/lint v0.0.0-20190930215403-16217165b5de
	golang.org/x/tools v0.0.0-20191119175705-11e13f1c3fd7
	honnef.co/go/tools v0.0.1-2019.2.3
)
