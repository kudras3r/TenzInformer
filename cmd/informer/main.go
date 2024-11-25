package main

import (
	"flag"
	"log"

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

	if len(flag.Args()) > 2 {
		log.Fatal("To enough CL arguments, needed only: 'log', 'conf'")
	}

	// TODO ->
	// init logger
	logger, err := logger.NewLogger(logger.DEBUG, *logFile)
	if err != nil {
		log.Fatal("Cant init logger, check: "+*logFile, err)
	}
	_ = logger
	_ = confFile

	// init grabber
	// run sender
}
