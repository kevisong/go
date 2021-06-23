package flagx

import (
	"flag"
	"fmt"
	"os"
)

var version = "1.0.0"

func useFlag() {

	var (
		fConfigFile = flag.String(
			"config",
			"./config.yml",
			"path to config file, e.g. --config=./config.yml, default is ./config.yml",
		)
		fVersion = flag.Bool(
			"version",
			false,
			"prints version and exits",
		)
	)

	flag.Parse()

	if *fVersion {
		fmt.Printf("Version %s\n", version)
		os.Exit(0)
	}

	if fConfigFile != nil {
		// Init Configuration
	}

}
