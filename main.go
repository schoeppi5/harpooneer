package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/schoeppi5/harpooneer/config"
	"github.com/schoeppi5/harpooneer/logging"

	"github.com/sirupsen/logrus"
)

var (
	yamlPath = flag.String("config", "harpooneer.yaml", "Path to the config file")
)

func main() {
	flag.Parse()
	// load config
	conf := config.Load(*yamlPath)
	// init logger
	level, err := logrus.ParseLevel(conf.Logging.Level)
	if err != nil {
		log.Printf("Unable to parse level %s. Using level info", conf.Logging.Level)
		level = logrus.InfoLevel
	}
	log := logging.NewLogrus(
		level,
		os.Stdout,
		os.Stderr,
		ioutil.Discard,
	)
}
