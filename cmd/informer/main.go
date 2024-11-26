package main

import (
	"flag"
	"log"

	"github.com/kudras3r/TenzInfromer/internal/info/grab"
	"github.com/kudras3r/TenzInfromer/internal/lib/logger"
)

const (
	defLogFile  = "/var/log/tenzir-example/tenzinformer.log"
	defConfFile = "/etc/tenzir-example/config.yml"
)

func main() {
	var logFile, confFile *string

	// parse flags
	logFile = flag.String("log", defLogFile, "")
	confFile = flag.String("conf", defConfFile, "")

	flag.Parse()

	// TODO ->
	// init logger
	logger, err := logger.NewLogger(logger.DEBUG, *logFile)
	if err != nil {
		log.Fatal("Cant init logger, check: "+*logFile, err)
	}

	logger.INFO("running app")
	logger.INFO("confFile: " + *confFile)

	// grab info
	PCInfo, err := grab.PCInfo(*confFile)
	if err != nil {
		logger.ERROR("сant grab information: " + err.Error())
	} else {
		logger.INFO("info grabbed")
	}
	_ = PCInfo

	// run sender
}
