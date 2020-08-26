# Changelog

## Unreleased

## v0.9.0

# 🎉 OpenTelemetry Collector Contrib v0.9.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.9.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/about/) for deployment and configuration information.

## 🛑 Breaking changes 🛑
- Remove deprecated `lightstep` exporter (#828)

## 🚀 New components 🚀
- `statsd` receiver for ingesting StatsD messages (#566)

## 💡 Enhancements 💡
- `signalfx` exporter
   - Add disk usage translations (#760)
   - Add disk utilization translations (#782)
   - Add translation rule to drop redundant metrics (#809)
- `kubeletstats` receiver
  - Sync available volume metadata from /pods endpoint (#690)
  - Add ability to collect detailed data from PVC (#743)
- `awsxray` exporter: Translate SDK name/version into xray model (#755)
- `elastic` exporter: Translate semantic conventions to Elastic destination fields (#671)
- `stackdriver` exporter: Add point count metric (#757)
- `awsxray` receiver
  - Ported the TCP proxy from the X-Ray daemon (#774)
  - Convert to OTEL trace format (#691)

## 🧰 Bug fixes 🧰
- `kubeletstats` receiver: Do not break down metrics batch (#754)
- `host` observer: Fix issue on darwin where ports listening on all interfaces are not correctly accounted for (#582)
- `newrelic` exporter: Fix panic on missing span status (#775)

## v0.8.0

# 🎉 OpenTelemetry Collector Contrib v0.8.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.8.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/about/) for deployment and configuration information.

## 🚀 New components 🚀

- Receivers
  - `prometheusexec` subprocess manager (##499)

## 💡 Enhancements 💡

- `signalfx` exporter
  - Add/Update metric translations (#579, #584, #639, #640, #652, #662)
  - Add support for calculate new metric translator (#644)
  - Add renaming rules for load metrics (#664)
  - Update `container.name` to `k8s.container.name` in default translation rule (#683)
  - Rename working-set and page-fault metrics (#679)
- `awsxray` exporter
  - Translate exception event into xray exception (#577)
  - Add ingestion of X-Ray segments via UDP (#502)
  - Parse Java stacktrace and populate in xray cause (#687)
- `kubeletstats` receiver
  - Add metric_groups option (#648)
  - Set datapoint timestamp in receiver (#661)
  - Change `container.name` label to `k8s.container.name` (#680)
  - Add working-set and page-fault metrics (#666)
  - Add basic support for volume metrics (#667)
- `stackdriver` trace exporter: Move to new interface and pdata (#486)
- `metricstranform` processor: Keep timeseries and points in order after aggregation (#663)
- `k8scluster` receiver: Change `container.spec.name` label to `k8s.container.name` (#681)
- Migrate receiver creator to internal data model (#701)
- Add ec2 support to `resourcedetection` processor (#587)
- Enable timeout, sending queue and retry for SAPM exporter (#707)

## 🧰 Bug fixes 🧰

- `azuremonitor` exporter: Correct HTTP status code success mapping (#588)
- `k8scluster` receiver: Fix owner reference in metadata updates (#649)
- `awsxray` exporter: Fix handling of db system (#697)

## 🚀 New components 🚀

- Skeleton for AWS ECS container metrics receiver (#463)
- `prometheus_exec` receiver (#655)

## v0.7.0

# 🎉 OpenTelemetry Collector Contrib v0.7.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.7.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/about/) for deployment and configuration information.

## 🛑 Breaking changes 🛑

- `awsxray` receiver updated to support udp: `tcp_endpoint` config option renamed to `endpoint` (#497)
- TLS config changed for `sapmreceiver` (#488) and `signalfxreceiver` receivers (#488)

## 🚀 New components 🚀

- Exporters
  - `sentry` adds tracing exporter for [Sentry](https://sentry.io/) (#565)
- Extensions
  - `endpoints` observer: adds generic endpoint watcher (#427)
  - `host` observer: looks for listening network endpoints on host (#432)

## 💡 Enhancements 💡

- Update `honeycomb` exporter for v0.8.0 compatibility
- Extend `metricstransform` processor to be able to add a label to an existing metric (#441)
- Update `kubeletstats` metrics according to semantic conventions (#475)
- Updated `awsxray` receiver config to use udp (#497)
- Add `/pods` endpoint support in `kubeletstats` receiver to add extra labels (#569)
- Add metric translation options to `signalfx` exporter (#477, #501, #571, #573)

## 🧰 Bug fixes 🧰

- `azuremonitor` exporter: Mark spanToEnvelope errors as permanent (#500)

## v0.6.0

# 🎉 OpenTelemetry Collector Contrib v0.6.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.6.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/about/) for deployment and configuration information.

## 🛑 Breaking changes 🛑

- Removed `jaegarlegacy` (#397) and `zipkinscribe` receivers (#410)
- `kubeletstats` receiver: Renamed `k8s.pod.namespace` pod label to `k8s.namespace.name` and `k8s.container.name` container label to `container.name`

## 🚀 New components 🚀

- Processors
  - `metricstransform` renames/aggregates within individual metrics (#376) and allow changing the data type between int and float (#402)

## 💡 Enhancements 💡

- `awsxray` exporter: Use `peer.service` as segment name when set. (#385)
- `splunk` exporter: Add trace exports support (#359, #399)
- Build and publish Windows MSI (#408) and DEB/RPM Linux packages (#405)

## 🧰 Bug fixes 🧰

- `kubeletstats` receiver:
  - Fixed NPE for newly created pods (#404)
  - Updated to latest change in the ReceiverFactoryOld interface (#401)
  - Fixed logging and self reported metrics (#357)
- `awsxray` exporter: Only convert SQL information for SQL databases. (#379)
- `resourcedetection` processor: Correctly obtain machine-type info from gce metadata (#395)
- `k8scluster` receiver: Fix container resource metrics (#416)

## v0.5.0

Released 01-07-2020

# 🎉 OpenTelemetry Collector Contrib v0.5.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.5.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/about/) for deployment and configuration information.

## 🚀 New components 🚀

- Processors
  - `resourcedetection` to automatically detect the resource based on the configured set of detectors (#309)

## 💡 Enhancements 💡

- `kubeletstats` receiver: Support for ServiceAccount authentication (#324)
- `signalfx` exporter and receiver
  - Add SignalFx metric token passthrough and config option (#325)
  - Set default endpoint of `signalfx` receiver to `:9943` (#351)
- `awsxray` exporter: Support aws plugins EC2/ECS/Beanstalk (#343)
- `sapm` exporter and receiver: Add SAPM access token passthrough and config option (#349)
- `k8s` processor: Add metrics support (#358)
- `k8s` observer: Separate annotations from labels in discovered pods (#363)

## 🧰 Bug fixes 🧰

- `honeycomb` exporter: Remove shared use of libhoney from goroutines (#305)

## v0.4.0

Released 17-06-2020

# 🎉 OpenTelemetry Collector Contrib v0.4.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.4.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/about/) for deployment and configuration information.

## 🛑 Breaking changes 🛑

  - `signalfx` exporter `url` parameter changed to `ingest_url` (no impact if only using `realm` setting)

## 🚀 New components 🚀

- Receivers
  - `receiver_creator` to create receivers at runtime (#145), add observer support to receiver_creator (#173), add rules support (#207), add dynamic configuration values (#235) 
  - `kubeletstats` receiver (#237) 
  - `prometheus_simple` receiver (#184) 
  - `kubernetes-cluster` receiver (#175) 
  - `redis` receiver (#138)
- Exporters
  - `alibabacloudlogservice` exporter (#259) 
  - `SplunkHEC` metrics exporter (#246)
  - `elastic` APM exporter (#240)
  - `newrelic` exporter (#229) 
- Extensions
  - `k8s` observer (#185) 

## 💡 Enhancements 💡

- `awsxray` exporter
  - Use X-Ray convention of segment name == service name (#282)
  - Tweak xray export to improve rendering of traces and improve parity (#241)
  - Add handling for spans received with nil attributes (#212)
- `honeycomb` exporter
  - Use SendPresampled (#291)
  - Add span attributes as honeycomb event fields (#271)
  - Support resource labels in Honeycomb exporter (#20)
- `k8s` processor
  - Add support of Pod UID extraction to k8sprocessor (#219)
  - Use `k8s.pod.ip` to record resource IP instead of just `ip` (#183)
  - Support same authentication mechanism as other kubernetes components do (#307)
- `sapm` exporter: Add TLS for SAPM and SignalFx receiver (#215)
- `signalfx` exporter
  - Add metric metadata syncer to SignalFx exporter (#231)
  - Add TLS for SAPM and SignalFx receiver (#215)
- `stackdriver` exporter: Add support for resource mapping in config (#163)

## 🧰 Bug fixes 🧰

- `awsxray` exporter: Wrap bad request errors for proper handling by retry queue (#205)
- `lightstep` exporter: Ensure Lightstep exporter doesnt crash on nil node (#250)
- `sapm` exporter: Do not break Jaeger traces before sending downstream (#193)
- `k8s` processor: Ensure Jaeger spans work in passthrough mode (262)

## 🧩 Components 🧩

### Receivers

| Traces | Metrics |
|:-------:|:-------:|
| Jaeger Legacy | Carbon |
| SAPM (SignalFx APM) | Collectd | 
| Zipkin Scribe | K8s Cluster |
| | Redis |
| |  SignalFx | 
| | Simple Prometheus |
| | Wavefront |

### Processors

- K8s

### Exporters

| Commercial | Community |
|:------------:|:-----------:|
| Alibaba Cloud Log Service | Carbon |
| AWS X-ray | Elastic |
| Azure Monitor | Jaeger Thrift |
| Honeycomb | Kinesis |
| Lightstep |
| New Relic |
| SAPM (SignalFx APM) | 
| SignalFx (Metrics) |
| Splunk HEC |
| Stackdriver (Google) |

### Extensions

- Observer
  - K8s

## v0.3.0 Beta

Released 2020-03-30

### Breaking changes

-  Make prometheus receiver config loading strict. #697 
Prometheus receiver will now fail fast if the config contains unused keys in it.

### Changes and fixes

- Enable best effort serve by default of Prometheus Exporter (https://github.com/orijtech/prometheus-go-metrics-exporter/pull/6)
- Fix null pointer exception in the logging exporter #743 
- Remove unnecessary condition to have at least one processor #744 
- Updated Honeycomb exported to `honeycombio/opentelemetry-exporter-go v0.3.1`

### Features

Receivers / Exporters:

* AWS X-Ray
* Carbon
* CollectD
* Honeycomb
* Jaeger
* Kinesis
* LightStep
* OpenCensus
* OpenTelemetry
* SAPM
* SignalFx
* Stackdriver
* Wavefront
* Zipkin
* Zipkin Scribe


Processors:

* Attributes
* Batch
* Memory Limiter
* Queued Retry
* Resource
* Sampling
* Span
* Kubernetes

Extensions:

* Health Check
* Performance Profiler
* zPages


## v0.2.8

Released 2020-03-25

Alpha v0.2.8 of OpenTelemetry Collector Contrib.

- Implemented OTLP receiver and exporter.
- Added ability to pass config to the service programmatically (useful for custom builds).
- Improved own metrics / observability.


## v0.2.7

Released 2020-03-17

### Self-Observability
- New command-line switch to control legacy and new metrics. Users are encouraged
to experiment and migrate to the new metrics.
- Improved error handling on shutdown.


### Processors
- Fixed passthrough mode k8sprocessor.
- Added `HASH` action to attribute processor.

### Receivers and Exporters
- Added Honeycomb exporter.
- Added LightStep exporter.
- Added regular expression for Carbon receiver, allowing the metric name to be broken into proper label keys and values.
- Updated Stackdriver exporter to use a new batch API.


## v0.2.6 Alpha

Released 2020-02-18

### Self-Observability
- Updated metrics prefix to `otelcol` and expose command line argument to modify the prefix value.
- Batch dropped span now emits zero when no spans are dropped.

### Processors
- Extended Span processor to have include/exclude span logic.
- Ability to choose strict or regexp matching for include/exclude filters.

### Receivers and Exporters
- Added Carbon receiver and exporter.
- Added Wavefront receiver.


## v0.0.5 Alpha

Released 2020-01-30

- Regexp-based filtering of span names.
- Ability to extract attributes from span names and rename span.
- File exporter for debugging.
- Span processor is now enabled by default.

## v0.0.1 Alpha

Released 2020-01-11

First release of OpenTelemetry Collector Contrib.


[v0.3.0]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.2.8...v0.3.0
[v0.2.8]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.2.7...v0.2.8
[v0.2.7]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.2.6...v0.2.7
[v0.2.6]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.0.5...v0.2.6
[v0.0.5]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.0.1...v0.0.5
[v0.0.1]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.0.1