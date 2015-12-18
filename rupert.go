package rupert

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/leighmacdonald/mika"
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
	startTime int32
)

func Start() {
	db := initDb()
	defer db.Db.Close()

	mika.SetupLogger("info", true)
	cfg, err := readConfig(*config_file)
	if err != nil {
		log.Fatal(err.Error())
	}
	setConfig(cfg)
	ListenAndServe(NewEngine())
}

func init() {
	startTime = int32(time.Now().Unix())
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
}
