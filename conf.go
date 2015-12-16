package rupert

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"sync"
)

var (
	// Global config instance & lock
	config     *Configuration
	configLock = new(sync.RWMutex)
)

type Configuration struct {
	// Enabled debug functions
	Debug bool

	// Enable testing mode which bypasses some functionality
	// Do not set this to true in production ever!
	Testing bool

	// A loglevel to use "info", "error", "warn", "debug", "fatal", "panic"
	LogLevel string

	// URI for the tracker listen host :34000
	ListenHost string

	// Redis hostname
	RedisHost string

	// Redis password, empty string for none
	RedisPass string

	// Maximum amount of idle redis connection to allow to idle
	RedisMaxIdle int

	// Redis database number to use
	RedisDB int

	// Path to the SSL private key
	SSLPrivateKey string

	// Path to the SSL CA cert
	SSLCert string

	// Use colours log output
	ColourLogs bool
}

// LoadConfig reads in a json based config file from the path provided and updated
// the currently active application configuration
func readConfig(config_file string) (*Configuration, error) {

	cfg := new(Configuration)

	log.WithFields(log.Fields{
		"config_file": config_file,
	}).Info("Loading config")
	log.Println(config_file)
	file, err := ioutil.ReadFile(config_file)
	if err != nil {
		log.WithFields(log.Fields{
			"fn":          "LoadConfig",
			"err":         err.Error(),
			"config_file": config_file,
		}).Fatal("Failed to open config file")
		return nil, err
	}

	if err = json.Unmarshal(file, cfg); err != nil {
		log.WithFields(log.Fields{
			"fn":          "LoadConfig",
			"err":         err.Error(),
			"config_file": config_file,
		}).Error("Failed to parse config file, cannot continue")
		return cfg, err
	}
	return cfg, nil
}

func setConfig(cfg *Configuration) {
	configLock.Lock()
	config = cfg
	configLock.Unlock()
}
