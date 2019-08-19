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
}

// MetricsOpts provides the configuration for the metrics endpoint
type MetricsOpts struct {
	Enabled          bool   `yaml:"enabled"`
	BasicAuthEnabled bool   `yaml:"basicAuthEnabled"`
	Port             int    `yaml:"port"`
	Username         string `yaml:"basicAuthUsername"`
	Password         string `yaml:"-"`
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
			Port:            80,
			ResponseCode:    200,
			ResponseMessage: "OK",
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
		return err
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
	if o.HeartbeatOpts.Port == o.MetricsOpts.Port {
		log.Errorf("Hearbeat endpoint and metrics endpoint port can not be equal. Current Port for both is %d", o.HeartbeatOpts.Port)
		isInValid = true
	}

	if o.MetricsOpts.BasicAuthEnabled {
		if o.Username == "" {
			log.Error("Username for metrics basic auth endpoint is empty. Please set an username")
			isInValid = true
		}

		value, exists := os.LookupEnv(envBasicAuthUsername)
		if !exists {
			log.Errorf("Could not lookup \"%s\" in environment variables. Please exportdeveloperyour preferred username", envBasicAuthUsername)
			isInValid = true
		} else if value == "" {
			log.Errorf("Environemt variable \"%s\" for metrics basic auth is empty. Please export your preferred username", envBasicAuthUsername)
			isInValid = true
		}

	}

	if isInValid {
		return Error{
			fmt.Sprint("error: opts package: Your configuration is not valid. Please check it."),
		}
	}

	return nil
}
