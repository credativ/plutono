+++
title = "Jaeger instrumentation"
description = "Jaeger traces emitted and propagation by Plutono"
keywords = ["plutono", "jaeger", "tracing"]
weight = 900
+++

# Jaeger instrumentation

Plutono supports [Jaeger tracing](https://www.jaegertracing.io/).

Plutono can emit Jaeger traces for its HTTP API endpoints and propagate Jaeger trace information to data sources.
All HTTP endpoints are logged evenly (annotations, dashboard, tags, and so on).
When a trace ID is propagated, it is reported with operation 'HTTP /datasources/proxy/:id/*'.

Refer to [Configuration]({{< relref "configuration.md#tracing-jaeger" >}}) for information about enabling Jaeger tracing.
