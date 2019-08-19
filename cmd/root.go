package cmd

import (
	"os"

	"github.com/fwiedmann/heartbeat/pkg/opts"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:          "Heartbeat",
	Short:        "Hearbeat is a simple health endpoint with additional prometheus metrics",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		o := opts.New(configFile, logLevel)
		if err := o.InitOpts(); err != nil {
			return err
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
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
