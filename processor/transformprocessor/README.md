# Transform Processor

| Status                   |                       |
| ------------------------ | --------------------- |
| Stability                | [alpha]               |
| Supported pipeline types | traces, metrics, logs |
| Distributions            | [contrib]             |

The transform processor modifies telemetry based on configuration using the [Telemetry Query Language](https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/processing.md#telemetry-query-language).
It takes a list of queries which are performed in the order specified in the config.

Queries are composed of the following parts
- Path expressions: Fields within the incoming data can be referenced using expressions composed of the names as defined
in the OTLP protobuf definition. e.g., `status.code`, `attributes["http.method"]`. If the path expression begins with
`resource.` or `instrumentation_library.`, it will reference those values.  For metrics, `name`, `description`, `unit`, `type`, `is_monotonic`, and `aggregation_temporality` are accessed via `metric.`
  - The name `instrumentation_library` within OpenTelemetry is currently under discussion and may be changed in the future.
  - Metric data types are `None`, `Gauge`, `Sum`, `Histogram`, `ExponentialHistogram`, and `Summary`
  - `aggregation_temporality` is converted to and from the [protobuf's numeric definition](https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/metrics/v1/metrics.proto#L291).  Interact with this field using 0, 1, or 2.
  - Until the grammar can handle booleans, `is_monotic` is handled via strings the strings `"true"` and `"false"`.
- Literals: Strings, ints, and floats can be referenced as literal values.  Byte slices can be references as a literal value via a hex string prefaced with `0x`, such as `0x0001`. 
- Function invocations: Functions can be invoked with arguments matching the function's expected arguments
- Where clause: Telemetry to modify can be filtered by appending `where a <op> b`, with `a` and `b` being any of the above.

Supported functions:
- `SpanID(bytes)` - `bytes` is a byte slice of exactly 8 bytes. The function returns a SpanID from `bytes`. e.g., `SpanID(0x0000000000000000)`

- `TraceID(bytes)` - `bytes` is a byte slice of exactly 16 bytes. The function returns a TraceID from `bytes`. e.g., `TraceID(0x00000000000000000000000000000000)`

- `set(target, value)` - `target` is a path expression to a telemetry field to set `value` into. `value` is any value type.
e.g., `set(attributes["http.path"], "/foo")`, `set(name, attributes["http.route"])`. If `value` resolves to `nil`, e.g.
it references an unset map value, there will be no action.

- `keep_keys(target, string...)` - `target` is a path expression to a map type field. The map will be mutated to only contain
the fields specified by the list of strings. e.g., `keep_keys(attributes, "http.method")`, `keep_keys(attributes, "http.method", "http.route")`

- `truncate_all(target, limit)` - `target` is a path expression to a map type field. `limit` is a non-negative integer.  The map will be mutated such that all string values are truncated to the limit. e.g., `truncate_all(attributes, 100)` will truncate all string values in `attributes` such that all string values have less than or equal to 100 characters.  Non-string values are ignored.

- `limit(target, limit)` - `target` is a path expression to a map type field. `limit` is a non-negative integer.  The map will be mutated such that the number of items does not exceed the limit. e.g., `limit(attributes, 100)` will limit `attributes` to no more than 100 items. Which items are dropped is random.

- `replace_match(target, pattern, replacement)` - `target` is a path expression to a telemetry field, `pattern` is a string following [filepath.Match syntax](https://pkg.go.dev/path/filepath#Match), and `replacement` is a string. If `target` matches `pattern` it will get replaced with `replacement`. e.g., `replace_match(attributes["http.target"], "/user/*/list/*", "/user/{userId}/list/{listId}")`

- `replace_all_matches(target, pattern, replacement)` - `target` is a path expression to a map type field, `pattern` is a string following [filepath.Match syntax](https://pkg.go.dev/path/filepath#Match), and `replacement` is a string. Each string value in `target` that matches `pattern` will get replaced with `replacement`. e.g., `replace_all_matches(attributes, "/user/*/list/*", "/user/{userId}/list/{listId}")`

Metric only functions:
- `convert_sum_to_gauge()` - Converts incoming metrics of type "Sum" to type "Gauge", retaining the metric's datapoints. Noop for metrics that are not of type "Sum". 
**NOTE:** This function may cause a metric to break semantics for [Gauge metrics](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/datamodel.md#gauge). Use at your own risk.

- `convert_gauge_to_sum(aggregation_temporality, is_monotonic)` - `aggregation_temporality` specifies the resultant metric's aggregation temporality. `aggregation_temporality` may be `"cumulative"` or `"delta"`. `is_monotonic` specifies the resultant metric's monotonicity. `is_monotonic` is a boolean. Converts incoming metrics of type "Gauge" to type "Sum", retaining the metric's datapoints and setting its aggregation temporality and monotonicity accordingly. Noop for metrics that are not of type "Gauge". 
**NOTE:** This function may cause a metric to break semantics for [Sum metrics](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/datamodel.md#sums). Use at your own risk.

Supported where operations:
- `==` - matches telemetry where the values are equal to each other
- `!=` - matches telemetry where the values are not equal to each other

Example configuration:
```yaml
receivers:
  otlp:
    protocols:
      grpc:

exporters:
  nop

processors:
  transform:
    traces:
      queries:
        - set(status.code, 1) where attributes["http.path"] == "/health"
        - keep_keys(resource.attributes, "service.name", "service.namespace", "cloud.region")
        - set(name, attributes["http.route"])
        - replace_match(attributes["http.target"], "/user/*/list/*", "/user/{userId}/list/{listId}")
        - limit(attributes, 100)
        - limit(resource.attributes, 100)
        - truncate_all(attributes, 4096)
        - truncate_all(resource.attributes, 4096)
    metrics:
      queries:
        - set(metric.description, "Sum") where metric.type == "Sum"
        - keep_keys(resource.attributes, "host.name")
        - limit(attributes, 100)
        - truncate_all(attributes, 4096)
        - truncate_all(resource.attributes, 4096)
        - convert_sum_to_gauge() where metric.name == "system.processes.count"
        - convert_gauge_to_sum("cumulative", false) where metric.name == "prometheus_metric"
    logs:
      queries:
        - set(severity_text, "FAIL") where body == "request failed"
        - replace_all_matches(attributes, "/user/*/list/*", "/user/{userId}/list/{listId}")
        - set(body, attributes["http.route"])
        - keep_keys(resource.attributes, "service.name", "service.namespace", "cloud.region")
service:
  pipelines:
    logs:
      receivers: [otlp]
      processors: [transform]
      exporters: [nop]
    traces:
      receivers: [otlp]
      processors: [transform]
      exporters: [nop]
```

This processor will perform the operations in order for 

All spans

1) Set status code to OK for all spans with a path `/health`
2) Keep only `service.name`, `service.namespace`, `cloud.region` resource attributes
3) Set `name` to the `http.route` attribute if it is set
2) Replace the value of an attribute named `http.target` with `/user/{userId}/list/{listId}` if the value matched `/user/*/list/*`
4) Limit all span attributes such that each span has no more than 100 attributes.
5) Limit all resource attributes such that each resource no more than 100 attributes.
6) Truncate all span attributes such that no string value has more than 4096 characters.
7) Truncate all resource attributes such that no string value has more than 4096 characters.

All metrics and their data points

1) Set metric description to "Sum" if the metric type is "Sum"
2) Keep only the `host.name` resource attributes
4) Limit all data point attributes such that each data point has no more than 100 attributes.
6) Truncate all data point attributes such that no string value has more than 4096 characters.
7) Truncate all resource attributes such that no string value has more than 4096 characters.
8) Convert all metrics with name `system.processes.count` from a Sum to Gauge.
9) Convert all metrics with name `prometheus_metric` from Gauge to a cumulative, non-monotonic Sum.

All logs

1) Set severity text to FAIL if the body contains a string text "request failed"
2) Replace any attribute value that matches `/user/*/list/*` with `/user/{userId}/list/{listId}`
3) Set `body` to the `http.route` attribute if it is set
4) Keep only `service.name`, `service.namespace`, `cloud.region` resource attributes

[alpha]: https://github.com/open-telemetry/opentelemetry-collector#alpha
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
