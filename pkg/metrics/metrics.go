package metrics

import (
	"fmt"
	"net/http"

	"github.com/fwiedmann/heartbeat/pkg/opts"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	registry = prometheus.NewRegistry()

	// HeartbeatVersion prom metric
	HeartbeatVersion = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heartbeat_version",
		Help: "Counter of request by each requester ",
	},
		[]string{"version"},
	)

	// HeartbeatRequester prom metric
	HeartbeatRequester = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heartbeat_requester",
		Help: "Counter of request by each requester ",
	},
		[]string{"host", "method"},
	)
	// HeartbeatTotalRequests prom metric
	HeartbeatTotalRequests = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heartbeat_total_requests",
		Help: "Total Number of requests.",
	},
		[]string{"method"},
	)
)

func init() {
	registry.MustRegister(HeartbeatVersion, HeartbeatRequester, HeartbeatTotalRequests)
}

// StartMetricsEndpoint starts metrics endpoint
func StartMetricsEndpoint(o opts.MetricsOpts) error {
	listenerPort := fmt.Sprintf(":%d", o.Port)
	http.Handle(o.Path, promhttp.Handler())
	if err := http.ListenAndServe(listenerPort, promhttp.HandlerFor(registry, promhttp.HandlerOpts{})); err != nil {
		return err
	}
	return nil
}
