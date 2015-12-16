package rupert

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"runtime"
	"time"
)

var (
	profile     = flag.String("profile", "", "write cpu profile to file")
	config_file = flag.String("config", "./config.json", "Config file path")
	num_procs   = flag.Int("procs", runtime.NumCPU(), "Number of CPU cores to use (default: ($num_cores-1))")

	// This is a special variable that is set by the go linker
	// If you do not build the project with make, or specify the linker settings
	// when building this will result in an empty version string
	Version = "master"

	// Timestamp of when the program first stared up
	StartTime int32
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

func main() {
	Start()
}

func Start() {
	db := initDb()
	defer db.Db.Close()

	initLogger("info", true)
	cfg, err := readConfig(*config_file)
	if err != nil {
		log.Fatal(err.Error())
	}
	setConfig(cfg)
	NewEngine().Run(":8081")
}

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
}
