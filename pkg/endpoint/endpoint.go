package endpoint

import (
	"fmt"
	"net/http"

	"github.com/fwiedmann/heartbeat/pkg/metrics"

	log "github.com/sirupsen/logrus"

	"github.com/fwiedmann/heartbeat/pkg/opts"
	"github.com/prometheus/client_golang/prometheus"
)

// StartHeartbeatEndpoint starts the heartbeat endpoint
func StartHeartbeatEndpoint(o opts.HeartbeatOpts) error {
	log.Infof("Starting heartbeat endpoint on port: \"%d\", path: \"%s\"", o.Port, o.Path)

	listenerPort := fmt.Sprintf(":%d", o.Port)
	heartbeatHandler := createHandler(o)

	http.HandleFunc(o.Path, heartbeatHandler)
	if err := http.ListenAndServe(listenerPort, nil); err != nil {
		return err
	}
	return nil

}

func createHandler(o opts.HeartbeatOpts) func(w http.ResponseWriter, r *http.Request) {
	log.Debugf("Created HeartbeatHandler with response code: \"%d\", response message: \"%s\"", o.ResponseCode, o.ResponseMessage)
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Incoming request: Host \"%s\", Method: \"%s\"  ", r.Host, r.Method)

		metrics.HeartbeatRequester.With(prometheus.Labels{"host": r.Host, "method": r.Method}).Inc()
		metrics.HeartbeatTotalRequests.With(prometheus.Labels{"method": r.Method}).Inc()

		w.WriteHeader(o.ResponseCode)
		if _, err := w.Write([]byte(o.ResponseMessage)); err != nil {
			log.Error(err)
		}

	}

}
