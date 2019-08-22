package opts

import (
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const envBasicAuthUsername = "HEARTBEAT_METRICS_PASSWORD"

// Error implements the error interface
type Error struct {
	err string
}

func (e Error) Error() string {
	return e.err
}

// HeartbeatOpts provides the configuration for the heartbeat endpoint
type HeartbeatOpts struct {
	Port            int    `yaml:"port"`
	ResponseCode    int    `yaml:"responseCode"`
	ResponseMessage string `yaml:"responseMessage"`
	Path            string `yaml:"path"`
}

// MetricsOpts provides the configuration for the metrics endpoint
type MetricsOpts struct {
	Enabled bool   `yaml:"enabled"`
	Port    int    `yaml:"port"`
	Path    string `yaml:"path"`
}

// Opts config
type Opts struct {
	HeartbeatOpts `yaml:"heartbeatEndpoint"`
	MetricsOpts   `yaml:"metricsEndpoint"`
	LogLevel      string `yaml:"-"`
	ConfigFile    string `yaml:"-"`
}

// New returns an empty Object struct
func New(config, logLevel string) *Opts {
	return &Opts{
		HeartbeatOpts: HeartbeatOpts{
			Port:            8080,
			ResponseCode:    200,
			ResponseMessage: "OK",
			Path:            "/heartbeat",
		},
		MetricsOpts: MetricsOpts{
			Port:    9100,
			Path:    "/metrics",
			Enabled: true,
		},
		LogLevel:   logLevel,
		ConfigFile: config,
	}
}

// InitOpts initialize hearbeats options from the configuration file
func (o *Opts) InitOpts() error {

	logLevel, err := log.ParseLevel(o.LogLevel)
	if err != nil {
		return err
	}

	log.SetLevel(logLevel)
	log.Debug("Check if file exists")
	if _, err = os.Open(o.ConfigFile); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		log.Info("No configuration was found. Using default values.")
		return nil
	}

	fileContent, err := ioutil.ReadFile(o.ConfigFile)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(fileContent, o); err != nil {
		return err
	}
	log.Debugf("Unmarshaled configfile: %s \n %+v", o.ConfigFile, *o)

	if err = o.validateOpts(); err != nil {
		return err
	}

	log.Debug("Configuration is valid")
	return nil

}

func (o *Opts) validateOpts() error {
	var isInValid bool

	o.HeartbeatOpts.Path = checkHandlerPath(o.HeartbeatOpts.Path)
	o.MetricsOpts.Path = checkHandlerPath(o.MetricsOpts.Path)

	if o.HeartbeatOpts.Port == o.MetricsOpts.Port {
		log.Errorf("Hearbeat endpoint and metrics endpoint port can not be equal. Current Port for both is %d", o.HeartbeatOpts.Port)
		isInValid = true
	}

	if isInValid {
		return Error{
			fmt.Sprint("error: opts package: Your configuration is not valid. Please check it."),
		}
	}

	return nil
}

func checkHandlerPath(path string) string {
	log.Debugf("Check if provided path \"%s\" is missing \"/\"", path)
	splitted := []byte(path)
	if splitted[0] != '/' {
		return fmt.Sprintf("/%s", path)
	}
	return path
}
