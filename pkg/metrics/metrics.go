package metrics

import (
	"fmt"
	"net/http"

	"github.com/fwiedmann/heartbeat/pkg/template"

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

	tmpl, err := template.New("Metrics endpoint", o.Path)
	if err != nil {
		return err
	}
	metricsSiteInfoHandler, err := tmpl.GetTempaltedHandler()
	if err != nil {
		return err
	}

	listenerPort := fmt.Sprintf(":%d", o.Port)

	server := http.NewServeMux()
	server.Handle(o.Path, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	if o.Path != "/" {
		server.Handle("/", metricsSiteInfoHandler)
	}

	if err := http.ListenAndServe(listenerPort, server); err != nil {
		return err
	}

	return nil
}
