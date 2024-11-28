/*
	There we grab system info from confFile and send it to tenzir.

	There are three main steps: grab - save - send data.

	First of all, we parse system info from defConfFile by def, but
	you can set another with --conf run flag.

	Next step is save data at tmp.json (in our case).
	But another formats may be added in the future in save pkg.

	Finally - send step.
	It works with running '| import' pipeline every 10 secs
	(check official tenzir docs).

	After we can take data from tenzir storage by running 'export |' pipeline.
*/

package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kudras3r/TenzInfromer/internal/info/grab"
	"github.com/kudras3r/TenzInfromer/internal/info/save"
	"github.com/kudras3r/TenzInfromer/internal/info/send"
	"github.com/kudras3r/TenzInfromer/internal/lib/logger"
)

// hardcode ! TODO
const (
	defLogFile  = "/var/log/tenzir-example/tenzinformer.log"
	defConfFile = "/etc/tenzir-example/config.yml"
	tmpJSONFile = "/home/kud/Code/go/src/TenzInformer/storage/tmp.json"
)

func main() {
	var logFile, confFile *string

	// parse flags
	logFile = flag.String("log", defLogFile, "")
	confFile = flag.String("conf", defConfFile, "")
	flag.Parse()

	// init logger
	logger, err := logger.NewLogger(logger.DEBUG, *logFile)
	if err != nil {
		log.Fatal("cannot init logger, check: "+*logFile, err)
	}

	logger.INFO("running app")
	logger.WARN("confFile: " + *confFile)

	// grab info
	PCInfo, err := grab.PCInfo(*confFile)
	if err != nil {
		logger.ERROR("cannot grab information: " + err.Error())
	} else {
		logger.INFO("info grabbed from: " + *confFile)
	}

	// serialize
	jsonData, err := json.MarshalIndent(PCInfo, "", " ")
	if err != nil {
		logger.ERROR("cannot convert to json: " + err.Error())
	} else {
		logger.INFO("convert data to json")
	}

	// save
	err = save.JSON(jsonData, tmpJSONFile)
	if err != nil {
		logger.ERROR("cannot save json: " + err.Error())
	} else {
		logger.INFO("save json to: " + tmpJSONFile)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// send data every 10 secs
	go func() {
		for {
			if err = send.JSON(tmpJSONFile); err != nil {
				logger.FATAL("cannot send json: " + err.Error())
			} else {
				logger.INFO("send data")
			}

			time.Sleep(10 * time.Second)
		}
	}()
	<-sigChan

	logger.INFO("shut down gracefully!")
}
