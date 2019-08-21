package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/fwiedmann/heartbeat/pkg/metrics"

	"github.com/fwiedmann/heartbeat/pkg/endpoint"
	"github.com/fwiedmann/heartbeat/pkg/opts"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:          "Heartbeat",
	Short:        "Hearbeat is a simple health endpoint with additional prometheus metrics",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ch := make(chan error, 2)

		o := opts.New(configFile, logLevel)
		if err := o.InitOpts(); err != nil {
			return err
		}

		go func() {
			ch <- endpoint.StartHeartbeatEndpoint(o.HeartbeatOpts)
		}()
		if o.MetricsOpts.Enabled {
			go func() {
				ch <- metrics.StartMetricsEndpoint(o.MetricsOpts)
			}()

		}

		for err := range ch {
			if err != nil {
				return err
			}
		}
		return nil
	},
}

var configFile string
var logLevel string

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "./config.yaml", "config file for endpoint configuration")
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "Set loglevel. Default is info")
}

// Execute executes the rootCmd
func Execute(version string) {

	metrics.HeartbeatVersion.With(prometheus.Labels{"version": version})

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
