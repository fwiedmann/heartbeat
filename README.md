<p align="center">
  <a href="https://example.com/">
    <img src="img/heartbeat_gopher.png" width=72 height=72>
  </a>

  <h3 align="center">Heartbeat</h3>

  <p align="center">
    Simple HTTP healthcheck endpoint with Prometheus metrics
    <br>
    <a href="https://github.com/fwiedmann/heartbeat/issues/new?template=bug.md">Report bug</a>
    ·
    <a href="https://github.com/fwiedmann/heartbeat/issues/new?template=feature.md&labels=feature">Request feature</a>
  </p>
</p>

## Table of contents

-   [Quick start](#quick-start)
-   [Status](#status)
-   [Custom Configuration](#custom-configuration)
    -   [Options description](#options-description)
-   [Build](#build)
-   [Copyright and license](#copyright-and-license)

## Quick start

Run Heartbeat as Docker container

```bash
do stuff here
```

## Status

-   [x] Implement basic logic
-   \[]Add Docker support
-   \[] Add CI Pipeline 

## Custom Configuration

Basic config file: 

```yaml
heartbeatEndpoint:
  port: 8080
  path: "/heartbeat"
  responseCode: 200
  responseMessage: "OK"
metricsEndpoint:
  enabled: true
  port: 9100
  path: "/metrics"
```

### Options description

| Endpoint  | Option           | Type   | Default vaule |
| --------- | ---------------- | ------ | ------------- |
| Heartbeat | path             | srting | "/heartbeat"  |
| Heartbeat | port             | int    | 8080            |
| Heartbeat | response message | srting | "OK"          |
| Heartbeat | response code    | int    | 200           |
| Metrics   | path             | srting | "/metrics"    |
| Metrics   | port             | int    | 9100          |
| Metrics   | enabled          | bool   | true          |

## Build

```bash
make build

# or
go build -o heartbeat -ldflags "-X main.HeartbeatVersion=<VERSION>"
```

## Copyright and license

[MIT License](https://reponame/blob/master/LICENSE).
