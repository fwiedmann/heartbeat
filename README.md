<p align="center">
  <a href="https://example.com/">
    <img src="img/heartbeat_gopher.png" width=72 height=72>
  </a>

  <h3 align="center">Heartbeat</h3>
  ![badge](https://action-badges.now.sh/fwiedmann/heartbeat)

  <p align="center">
    Simple HTTP healthcheck endpoint with Prometheus metrics
    <br>
    <a href="https://github.com/fwiedmann/heartbeat/releases/latest">Latest release </a>
    ·
        <a href="https://hub.docker.com/r/wiedmannfelix/heartbeat">Docker Hub </a>
    ·
    <a href="https://github.com/fwiedmann/heartbeat/issues/new?template=bug.md">Report bug</a>
    ·
    <a href="https://github.com/fwiedmann/heartbeat/issues/new?template=feature.md&labels=feature">Request feature</a>
  </p>
</p>


## Table of contents

-   [Quick start](#quick-start)
- [Avialable flags](#available-flags)
-   [Metrics](#metrics)
-   [Custom Configuration](#custom-configuration)
    -   [Options description](#options-description)
-   [Build](#build)
-   [Copyright and license](#copyright-and-license)

## Quick start

Deploy Heartbeat to your Kubernetes cluster

```bash
kubectl apply -f https://raw.githubusercontent.com/fwiedmann/heartbeat/master/k8s/namespace.yaml
kubectl apply -f https://raw.githubusercontent.com/fwiedmann/heartbeat/master/k8s/configMap.yaml
kubectl apply -f https://raw.githubusercontent.com/fwiedmann/heartbeat/master/k8s/deployment.yaml
```

Run Heartbeat as single Docker container with / without config.yaml

```bash
docker run --rm -d -p 8080:8080 -p 9100:9100 wiedmannfelix/heartbeat:latest

docker run --rm -d -p 8080:8080 -p 9100:9100 -v config.yaml:/code/config.yaml wiedmannfelix/heartbeat:latest
```

Run Heartbeat binary directly on your host

```bash
wget https://github.com/fwiedmann/heartbeat/releases/download/1.0.0/heartbeat

chmod +x heartbeat

./heartbeat &
```

## Available flags

```bash
    Usage:
      Heartbeat [flags]

    Flags:
          --config string     config file for endpoint configuration (default "./config.yaml")
      -h, --help              help for Heartbeat
          --loglevel string   Set loglevel. Default is info (default "info")
```

## Metrics

| Metric                                  | Type  | Description                            |
| --------------------------------------- | ----- | -------------------------------------- |
| heartbeat_version{version=""}           | Gauge | Prints the latest version of Heartbeat |
| heartbeat_requester{host="", method=""} | Gauge | Counter of request by each requester   |
| heartbeat_total_requests{method=""}     | Gauge | Total Number of requests               |

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
| Heartbeat | port             | int    | 8080          |
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
