package rupert

import (
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"time"
)

// SetupLogger will configure logrus to use our config
// force_colour will enable colour codes to be used even if there is no TTY detected
func initLogger(log_level string, force_colour bool) {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:    force_colour,
		DisableSorting: true,
	})
	switch log_level {
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func Main() {
	dbmap := initDb()
	configLock.Lock()
	Config = initConfig("config.json")
	configLock.Unlock()
	initLogger("info", true)

	NewEngine().Run(":8081")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
