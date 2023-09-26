# Strange attribute `app_id` in demo_server/request_counts metrics

In this demo, client send HTTP request to `/hello/:app_id` or `/bye/:app_id`, and reporting demo_client/request_counts counter metrics with attribute `app_id = ...`. Then server will report demo_server/request_counts counter metrics with attribute `app_id = ...` too.

And the app_id is a fixed list: `"1024", "568", "106", "1025", "197"`

It's really strage that in demo_client/request_counts everything is okay, but `app_id` would be strange string like "con", "247.c", "n.i" in demo_server/request_counts.

## How to test

```
docker compose up -d --build
```

## Example logs from otel-collector

```
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:06.208596348 +0000 UTC
Timestamp: 2023-09-26 12:57:08.208135317 +0000 UTC
Count: 1
Sum: 1704.050588
Min: 1704.050588
Max: 1704.050588
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:06.208599414 +0000 UTC
Timestamp: 2023-09-26 12:57:08.208143211 +0000 UTC
Value: 1
ResourceMetrics #2
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:06.483365128 +0000 UTC
Timestamp: 2023-09-26 12:57:08.482927931 +0000 UTC
Count: 1
Sum: 1649.867000
Min: 1649.867000
Max: 1649.867000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:06.483366994 +0000 UTC
Timestamp: 2023-09-26 12:57:08.482935526 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:06.48336792 +0000 UTC
Timestamp: 2023-09-26 12:57:08.482937299 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:08.482939068 +0000 UTC
Value: 1
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 2
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(197c)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:06.483397814 +0000 UTC
Timestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Value: 1
ResourceMetrics #3
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:08.208135317 +0000 UTC
Timestamp: 2023-09-26 12:57:10.208781014 +0000 UTC
Count: 1
Sum: 1211.490551
Min: 1211.490551
Max: 1211.490551
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:08.208143211 +0000 UTC
Timestamp: 2023-09-26 12:57:10.208791247 +0000 UTC
Value: 1
ResourceMetrics #4
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:08.482927931 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483412769 +0000 UTC
Count: 1
Sum: 1081.090000
Min: 1081.090000
Max: 1081.090000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:08.482935526 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483420478 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:08.482937299 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483422635 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483424239 +0000 UTC
Value: 1
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 2
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(247)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(247c)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:08.48294048 +0000 UTC
Timestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Value: 1
ResourceMetrics #5
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.208781014 +0000 UTC
Timestamp: 2023-09-26 12:57:12.20787546 +0000 UTC
Count: 1
Sum: 671.750750
Min: 671.750750
Max: 671.750750
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.208781014 +0000 UTC
Timestamp: 2023-09-26 12:57:12.20787546 +0000 UTC
Count: 1
Sum: 30.803748
Min: 30.803748
Max: 30.803748
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 1
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.208791247 +0000 UTC
Timestamp: 2023-09-26 12:57:12.207884956 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.208791247 +0000 UTC
Timestamp: 2023-09-26 12:57:12.207884956 +0000 UTC
Value: 1
ResourceMetrics #6
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(247c)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(7.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(7.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(247)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 2
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:10.483426141 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Value: 1
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:10.483412769 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482747501 +0000 UTC
Count: 2
Sum: 677.141000
Min: 20.605000
Max: 656.536000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 1
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:10.483420478 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482761023 +0000 UTC
Count: 2
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 2
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:10.483422635 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482762902 +0000 UTC
Count: 2
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 2
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:12.482764413 +0000 UTC
Value: 0
ResourceMetrics #7
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.20787546 +0000 UTC
Timestamp: 2023-09-26 12:57:14.208680279 +0000 UTC
Count: 1
Sum: 277.991351
Min: 277.991351
Max: 277.991351
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 1
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.20787546 +0000 UTC
Timestamp: 2023-09-26 12:57:14.208680279 +0000 UTC
Count: 1
Sum: 148.866438
Min: 148.866438
Max: 148.866438
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 1
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.207884956 +0000 UTC
Timestamp: 2023-09-26 12:57:14.208700185 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.207884956 +0000 UTC
Timestamp: 2023-09-26 12:57:14.208700185 +0000 UTC
Value: 1
ResourceMetrics #8
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:12.482747501 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482587542 +0000 UTC
Count: 2
Sum: 395.565000
Min: 132.758000
Max: 262.807000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 1
Buckets #8, Count: 1
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:12.482761023 +0000 UTC
Timestamp: 2023-09-26 12:57:14.48259518 +0000 UTC
Count: 2
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 2
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:12.482762902 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482597105 +0000 UTC
Count: 2
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 2
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482609443 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 2
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(5685)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(5685)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:12.482730885 +0000 UTC
Timestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Value: 1
        {"kind": "exporter", "data_type": "metrics", "name": "logging"}
2023-09-26T12:57:24.484Z        info    MetricsExporter {"kind": "exporter", "data_type": "metrics", "name": "logging", "resource metrics": 9, "metrics": 30, "data points": 103}
2023-09-26T12:57:24.485Z        info    ResourceMetrics #0
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482756018 +0000 UTC
Value: 1
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 2
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:14.482610733 +0000 UTC
Timestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Value: 1
ResourceMetrics #1
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.208832666 +0000 UTC
Timestamp: 2023-09-26 12:57:18.208388086 +0000 UTC
Count: 1
Sum: 1981.325110
Min: 1981.325110
Max: 1981.325110
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.208834732 +0000 UTC
Timestamp: 2023-09-26 12:57:18.208393197 +0000 UTC
Value: 1
ResourceMetrics #2
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:16.482742166 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483378857 +0000 UTC
Count: 1
Sum: 1951.117000
Min: 1951.117000
Max: 1951.117000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:16.482742166 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483378857 +0000 UTC
Count: 1
Sum: 136.771000
Min: 136.771000
Max: 136.771000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 1
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:16.482754221 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483387347 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:16.482754221 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483387347 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:16.482755399 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483390118 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:16.482755399 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483390118 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483392085 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(25i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(25ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 2
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(25ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(25ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:16.482758602 +0000 UTC
Timestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Value: 1
ResourceMetrics #3
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:18.208388086 +0000 UTC
Timestamp: 2023-09-26 12:57:20.208378329 +0000 UTC
Count: 2
Sum: 1027.708740
Min: 349.167982
Max: 678.540758
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 1
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:18.208393197 +0000 UTC
Timestamp: 2023-09-26 12:57:20.208384409 +0000 UTC
Value: 2
ResourceMetrics #4
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:18.483378857 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483422505 +0000 UTC
Count: 1
Sum: 616.385000
Min: 616.385000
Max: 616.385000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:18.483387347 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483429403 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:18.483390118 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483431268 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483432514 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(25ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(25ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(25ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(25i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(25i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(25ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 2
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:18.483393919 +0000 UTC
Timestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Value: 1
ResourceMetrics #5
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:20.208378329 +0000 UTC
Timestamp: 2023-09-26 12:57:22.208433726 +0000 UTC
Count: 1
Sum: 32.747684
Min: 32.747684
Max: 32.747684
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 1
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:20.208384409 +0000 UTC
Timestamp: 2023-09-26 12:57:22.208441162 +0000 UTC
Value: 1
ResourceMetrics #6
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:20.483422505 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482904016 +0000 UTC
Count: 1
Sum: 25.599000
Min: 25.599000
Max: 25.599000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 1
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:20.483429403 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482910921 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:20.483431268 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482912923 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482914404 +0000 UTC
Value: 1
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(568c)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 2
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(25ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(568c)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(568c)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(o/5)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(25i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:20.483433965 +0000 UTC
Timestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Value: 1
ResourceMetrics #7
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:22.208433726 +0000 UTC
Timestamp: 2023-09-26 12:57:24.208094615 +0000 UTC
Count: 1
Sum: 211.810865
Min: 211.810865
Max: 211.810865
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 1
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.208433726 +0000 UTC
Timestamp: 2023-09-26 12:57:24.208094615 +0000 UTC
Count: 1
Sum: 610.692230
Min: 610.692230
Max: 610.692230
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.208441162 +0000 UTC
Timestamp: 2023-09-26 12:57:24.208101628 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:22.208441162 +0000 UTC
Timestamp: 2023-09-26 12:57:24.208101628 +0000 UTC
Value: 1
ResourceMetrics #8
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:22.482904016 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483155524 +0000 UTC
Count: 1
Sum: 404.194000
Min: 404.194000
Max: 404.194000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 1
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:22.482904016 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483155524 +0000 UTC
Count: 1
Sum: 171.628000
Min: 171.628000
Max: 171.628000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 1
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:22.482910921 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483168084 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:22.482910921 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483168084 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:22.482912923 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483172173 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:22.482912923 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483172173 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:24.48317574 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 2
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(258c)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(258c)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(258)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(258)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(258c)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
        {"kind": "exporter", "data_type": "metrics", "name": "logging"}
2023-09-26T12:57:34.484Z        info    MetricsExporter {"kind": "exporter", "data_type": "metrics", "name": "logging", "resource metrics": 10, "metrics": 35, "data points": 115}
2023-09-26T12:57:34.485Z        info    ResourceMetrics #0
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.208094615 +0000 UTC
Timestamp: 2023-09-26 12:57:26.207935841 +0000 UTC
Count: 1
Sum: 422.103787
Min: 422.103787
Max: 422.103787
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 1
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.208101628 +0000 UTC
Timestamp: 2023-09-26 12:57:26.207943708 +0000 UTC
Value: 1
ResourceMetrics #1
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(2425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 2
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(2425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(2425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 2
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(242)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(242)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Value: 1
NumberDataPoints #15
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 1970-01-01 00:00:00 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:24.483155524 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482972103 +0000 UTC
Count: 1
Sum: 418.912000
Min: 418.912000
Max: 418.912000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 1
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:24.483155524 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482972103 +0000 UTC
Count: 1
Sum: 13.526000
Min: 13.526000
Max: 13.526000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 1
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:24.483168084 +0000 UTC
Timestamp: 2023-09-26 12:57:26.48297724 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:24.483168084 +0000 UTC
Timestamp: 2023-09-26 12:57:26.48297724 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:24.483172173 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482996077 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:24.483172173 +0000 UTC
Timestamp: 2023-09-26 12:57:26.482996077 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:26.483001039 +0000 UTC
Value: 0
ResourceMetrics #2
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.207935841 +0000 UTC
Timestamp: 2023-09-26 12:57:28.208783847 +0000 UTC
Count: 1
Sum: 27.617535
Min: 27.617535
Max: 27.617535
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 1
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.207935841 +0000 UTC
Timestamp: 2023-09-26 12:57:28.208783847 +0000 UTC
Count: 1
Sum: 559.520203
Min: 559.520203
Max: 559.520203
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.207943708 +0000 UTC
Timestamp: 2023-09-26 12:57:28.208792069 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.207943708 +0000 UTC
Timestamp: 2023-09-26 12:57:28.208792069 +0000 UTC
Value: 1
ResourceMetrics #3
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:26.482972103 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482507072 +0000 UTC
Count: 1
Sum: 476.341000
Min: 476.341000
Max: 476.341000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 1
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:26.48297724 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482514958 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:26.482996077 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482516816 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482518819 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 2
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 2
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:26.482958097 +0000 UTC
Timestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
ResourceMetrics #4
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:28.208783847 +0000 UTC
Timestamp: 2023-09-26 12:57:30.208864872 +0000 UTC
Count: 1
Sum: 146.674334
Min: 146.674334
Max: 146.674334
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 1
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:28.208792069 +0000 UTC
Timestamp: 2023-09-26 12:57:30.208871817 +0000 UTC
Value: 1
ResourceMetrics #5
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:28.482507072 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483545285 +0000 UTC
Count: 1
Sum: 85.568000
Min: 85.568000
Max: 85.568000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 1
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:28.482514958 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483553776 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:28.482516816 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483556103 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483557586 +0000 UTC
Value: 1
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 2
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 2
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:28.482520148 +0000 UTC
Timestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Value: 1
ResourceMetrics #6
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:30.208864872 +0000 UTC
Timestamp: 2023-09-26 12:57:32.208799626 +0000 UTC
Count: 1
Sum: 1357.430283
Min: 1357.430283
Max: 1357.430283
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:30.208871817 +0000 UTC
Timestamp: 2023-09-26 12:57:32.208814693 +0000 UTC
Value: 1
ResourceMetrics #7
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:30.483545285 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483162371 +0000 UTC
Count: 1
Sum: 1276.125000
Min: 1276.125000
Max: 1276.125000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:30.483553776 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483171322 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:30.483556103 +0000 UTC
Timestamp: 2023-09-26 12:57:32.48317346 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483174916 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 2
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(24i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 2
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:30.483558735 +0000 UTC
Timestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:22.482915687 +0000 UTC
Timestamp: 2023-09-26 12:57:24.483177626 +0000 UTC
Value: 1
ResourceMetrics #8
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.208799626 +0000 UTC
Timestamp: 2023-09-26 12:57:34.207928284 +0000 UTC
Count: 1
Sum: 1126.389236
Min: 1126.389236
Max: 1126.389236
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.208814693 +0000 UTC
Timestamp: 2023-09-26 12:57:34.207945019 +0000 UTC
Value: 1
ResourceMetrics #9
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:32.483162371 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483522356 +0000 UTC
Count: 1
Sum: 978.151000
Min: 978.151000
Max: 978.151000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:32.483171322 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483529937 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:32.48317346 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483531878 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483533703 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 2
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 2
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:32.483176284 +0000 UTC
Timestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Value: 1
        {"kind": "exporter", "data_type": "metrics", "name": "logging"}
2023-09-26T12:57:44.483Z        info    MetricsExporter {"kind": "exporter", "data_type": "metrics", "name": "logging", "resource metrics": 10, "metrics": 35, "data points": 118}
2023-09-26T12:57:44.484Z        info    ResourceMetrics #0
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:34.207928284 +0000 UTC
Timestamp: 2023-09-26 12:57:36.208738255 +0000 UTC
Count: 1
Sum: 709.450934
Min: 709.450934
Max: 709.450934
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:34.207945019 +0000 UTC
Timestamp: 2023-09-26 12:57:36.208745063 +0000 UTC
Value: 1
ResourceMetrics #1
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:34.483522356 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482607396 +0000 UTC
Count: 1
Sum: 704.392000
Min: 704.392000
Max: 704.392000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:34.483522356 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482607396 +0000 UTC
Count: 1
Sum: 24.581000
Min: 24.581000
Max: 24.581000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 1
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:34.483529937 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482616026 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:34.483529937 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482616026 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:34.483531878 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482618641 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:34.483531878 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482618641 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482620764 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 2
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 2
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(o/1)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(o/10)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
NumberDataPoints #16
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
ResourceMetrics #2
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.208738255 +0000 UTC
Timestamp: 2023-09-26 12:57:38.208280329 +0000 UTC
Count: 1
Sum: 26.659281
Min: 26.659281
Max: 26.659281
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 1
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
HistogramDataPoints #1
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.208738255 +0000 UTC
Timestamp: 2023-09-26 12:57:38.208280329 +0000 UTC
Count: 1
Sum: 630.095681
Min: 630.095681
Max: 630.095681
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1024)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.208745063 +0000 UTC
Timestamp: 2023-09-26 12:57:38.208288099 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.208745063 +0000 UTC
Timestamp: 2023-09-26 12:57:38.208288099 +0000 UTC
Value: 1
ResourceMetrics #3
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:36.482607396 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483259172 +0000 UTC
Count: 1
Sum: 565.006000
Min: 565.006000
Max: 565.006000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:36.482616026 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483266587 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:36.482618641 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483268931 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483270427 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 2
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 2
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Timestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Value: 1
ResourceMetrics #4
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.208280329 +0000 UTC
Timestamp: 2023-09-26 12:57:40.207927079 +0000 UTC
Count: 1
Sum: 159.785014
Min: 159.785014
Max: 159.785014
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 1
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.208288099 +0000 UTC
Timestamp: 2023-09-26 12:57:40.207933947 +0000 UTC
Value: 1
ResourceMetrics #5
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:38.483259172 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482588905 +0000 UTC
Count: 1
Sum: 121.657000
Min: 121.657000
Max: 121.657000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 1
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:38.483266587 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482596632 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:38.483268931 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482598646 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:40.48260024 +0000 UTC
Value: 1
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 2
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(706)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(1970)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 2
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(706)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:38.483271929 +0000 UTC
Timestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Value: 1
ResourceMetrics #6
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:40.207927079 +0000 UTC
Timestamp: 2023-09-26 12:57:42.208330418 +0000 UTC
Count: 1
Sum: 1768.896078
Min: 1768.896078
Max: 1768.896078
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:40.207933947 +0000 UTC
Timestamp: 2023-09-26 12:57:42.20833746 +0000 UTC
Value: 1
ResourceMetrics #7
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:40.482588905 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483348975 +0000 UTC
Count: 1
Sum: 1753.718000
Min: 1753.718000
Max: 1753.718000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 1
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:40.482596632 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483357837 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:40.482598646 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483360788 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483363702 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(197)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(1970)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 2
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(706)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(706)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:40.482602162 +0000 UTC
Timestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Value: 2
NumberDataPoints #16
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 1970-01-01 00:00:00 +0000 UTC
Value: 0
ResourceMetrics #8
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:42.208330418 +0000 UTC
Timestamp: 2023-09-26 12:57:44.208227858 +0000 UTC
Count: 1
Sum: 804.538858
Min: 804.538858
Max: 804.538858
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:42.20833746 +0000 UTC
Timestamp: 2023-09-26 12:57:44.208234611 +0000 UTC
Value: 1
ResourceMetrics #9
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:42.483348975 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482725287 +0000 UTC
Count: 1
Sum: 713.429000
Min: 713.429000
Max: 713.429000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:42.483357837 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482732912 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:42.483360788 +0000 UTC
Timestamp: 2023-09-26 12:57:44.48273498 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482736719 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 2
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 2
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
NumberDataPoints #16
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
        {"kind": "exporter", "data_type": "metrics", "name": "logging"}
2023-09-26T12:57:54.483Z        info    MetricsExporter {"kind": "exporter", "data_type": "metrics", "name": "logging", "resource metrics": 10, "metrics": 35, "data points": 115}
2023-09-26T12:57:54.484Z        info    ResourceMetrics #0
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.208227858 +0000 UTC
Timestamp: 2023-09-26 12:57:46.208692377 +0000 UTC
Count: 2
Sum: 554.719489
Min: 26.548645
Max: 528.170844
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 1
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.208234611 +0000 UTC
Timestamp: 2023-09-26 12:57:46.208699707 +0000 UTC
Value: 2
ResourceMetrics #1
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:44.482725287 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482628201 +0000 UTC
Count: 2
Sum: 488.404000
Min: 15.442000
Max: 472.962000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 1
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 1
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:44.482732912 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482643626 +0000 UTC
Count: 2
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 2
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:44.48273498 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482645858 +0000 UTC
Count: 2
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 2
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482647797 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 2
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 2
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(o/5)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(o/56)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Timestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Value: 1
NumberDataPoints #16
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:42.483367153 +0000 UTC
Timestamp: 2023-09-26 12:57:44.482738055 +0000 UTC
Value: 1
ResourceMetrics #2
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.208692377 +0000 UTC
Timestamp: 2023-09-26 12:57:48.208084289 +0000 UTC
Count: 1
Sum: 881.160301
Min: 881.160301
Max: 881.160301
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.208699707 +0000 UTC
Timestamp: 2023-09-26 12:57:48.208090306 +0000 UTC
Value: 1
ResourceMetrics #3
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:46.482628201 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482796747 +0000 UTC
Count: 1
Sum: 674.707000
Min: 674.707000
Max: 674.707000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:46.482643626 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482814068 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:46.482645858 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482815907 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482817621 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 2
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 2
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
NumberDataPoints #16
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:46.482649259 +0000 UTC
Timestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Value: 1
ResourceMetrics #4
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:48.208084289 +0000 UTC
Timestamp: 2023-09-26 12:57:50.208813957 +0000 UTC
Count: 1
Sum: 846.739005
Min: 846.739005
Max: 846.739005
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:48.208090306 +0000 UTC
Timestamp: 2023-09-26 12:57:50.208820851 +0000 UTC
Value: 1
ResourceMetrics #5
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:48.482796747 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482907685 +0000 UTC
Count: 1
Sum: 826.513000
Min: 826.513000
Max: 826.513000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:48.482814068 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482916148 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:48.482815907 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482917984 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:50.48292112 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 2
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(6.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(106.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 2
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(6.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:48.482818888 +0000 UTC
Timestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Value: 1
NumberDataPoints #16
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:34.483534856 +0000 UTC
Timestamp: 2023-09-26 12:57:36.482622701 +0000 UTC
Value: 1
ResourceMetrics #6
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:50.208813957 +0000 UTC
Timestamp: 2023-09-26 12:57:52.208764068 +0000 UTC
Count: 1
Sum: 203.983005
Min: 203.983005
Max: 203.983005
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 1
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:50.208820851 +0000 UTC
Timestamp: 2023-09-26 12:57:52.208770094 +0000 UTC
Value: 1
ResourceMetrics #7
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:50.482907685 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483496022 +0000 UTC
Count: 1
Sum: 97.746000
Min: 97.746000
Max: 97.746000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 1
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:50.482916148 +0000 UTC
Timestamp: 2023-09-26 12:57:52.48350384 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/bye/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:50.482917984 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483505958 +0000 UTC
Count: 1
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 1
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483507468 +0000 UTC
Value: 1
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(6.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 2
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(6.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(106.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(8425)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 2
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(568)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
NumberDataPoints #16
Data point attributes:
     -> app_id: Str(842)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:50.482922423 +0000 UTC
Timestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Value: 1
ResourceMetrics #8
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-client)
     -> process.command_args: Slice(["demo-client"])
     -> process.executable.name: Str(demo-client)
     -> process.executable.path: Str(/usr/local/bin/demo-client)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(client)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_client/request_latency
     -> Description: The latency of requests processed
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.208764068 +0000 UTC
Timestamp: 2023-09-26 12:57:54.207877635 +0000 UTC
Count: 1
Sum: 843.920913
Min: 843.920913
Max: 843.920913
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: demo_client/request_counts
     -> Description: The number of requests processed
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.208770094 +0000 UTC
Timestamp: 2023-09-26 12:57:54.207884138 +0000 UTC
Value: 1
ResourceMetrics #9
Resource SchemaURL: https://opentelemetry.io/schemas/1.21.0
Resource attributes:
     -> host.name: Str(demo-server)
     -> process.command_args: Slice(["demo-server"])
     -> process.executable.name: Str(demo-server)
     -> process.executable.path: Str(/usr/local/bin/demo-server)
     -> process.owner: Str(10001)
     -> process.pid: Int(1)
     -> process.runtime.description: Str(go version go1.21.1 linux/amd64)
     -> process.runtime.name: Str(go)
     -> process.runtime.version: Str(go1.21.1)
     -> service.name: Str(server)
     -> service.namespace: Str(demo)
     -> service.version: Str(0.1.0)
     -> telemetry.sdk.language: Str(go)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.18.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope demo-server 
Metric #0
Descriptor:
     -> Name: demo_server/request_counts
     -> Description: The number of requests received
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Delta
NumberDataPoints #0
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 2
NumberDataPoints #1
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #2
Data point attributes:
     -> app_id: Str(o/1)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #3
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #4
Data point attributes:
     -> app_id: Str(6.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #5
Data point attributes:
     -> app_id: Str(106)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #6
Data point attributes:
     -> app_id: Str(106.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #7
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #8
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #9
Data point attributes:
     -> app_id: Str(n.ic)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #10
Data point attributes:
     -> app_id: Str(con)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #11
Data point attributes:
     -> app_id: Str(102)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #12
Data point attributes:
     -> app_id: Str(n.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #13
Data point attributes:
     -> app_id: Str(6.i)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #14
Data point attributes:
     -> app_id: Str(con.)
     -> interface: Str(bye)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #15
Data point attributes:
     -> app_id: Str(1025)
     -> interface: Str(hello)
StartTimestamp: 2023-09-26 12:57:52.483508775 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482922446 +0000 UTC
Value: 1
NumberDataPoints #16
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 1970-01-01 00:00:00 +0000 UTC
Value: 0
ScopeMetrics #1
ScopeMetrics SchemaURL: 
InstrumentationScope github.com/gofiber/contrib/otelfiber 1.19.0
Metric #0
Descriptor:
     -> Name: http.server.duration
     -> Description: measures the duration inbound HTTP requests
     -> Unit: ms
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:52.483496022 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482936538 +0000 UTC
Count: 2
Sum: 1105.697000
Min: 364.946000
Max: 740.751000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 1
Buckets #9, Count: 1
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: http.server.request.size
     -> Description: measures the size of HTTP request messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:52.48350384 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482941204 +0000 UTC
Count: 2
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 2
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: http.server.response.size
     -> Description: measures the size of HTTP response messages
     -> Unit: By
     -> DataType: Histogram
     -> AggregationTemporality: Delta
HistogramDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.route: Str(/hello/:app_id)
     -> http.scheme: Str(http)
     -> http.status_code: Int(200)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:57:52.483505958 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482942784 +0000 UTC
Count: 2
Sum: 0.000000
Min: 0.000000
Max: 0.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 2
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: http.server.active_requests
     -> Description: measures the number of concurrent HTTP requests that are currently in-flight
     -> Unit: 1
     -> DataType: Sum
     -> IsMonotonic: false
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> http.flavor: Str(1.1)
     -> http.method: Str(GET)
     -> http.scheme: Str(http)
     -> net.host.name: Str(demo-server:2333)
StartTimestamp: 2023-09-26 12:55:46.483128655 +0000 UTC
Timestamp: 2023-09-26 12:57:54.482944075 +0000 UTC
Value: 0
        {"kind": "exporter", "data_type": "metrics", "name": "logging"}
```
