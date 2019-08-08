# heartbeat

Simple HTTP endpoint with Prometheus metrics

## To implement

- endpoints:
  - `/`: sumary of all endpoints
  - `/heartbeat`: response from config, send 200 OK
  - `/metrics`:
- configuration file:
  - port configuration for heartbeat endpoint and metrics
  - response body/message
  - basic auth for metric endpoint
- metrics:
  - show configuration
  - last request
  - list of all callers with: ip, method, url, request
