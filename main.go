package main

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/daddvted/fruitninja/fruitninja"
	log "github.com/sirupsen/logrus"
)

var appConfig fruitninja.FruitNinjaConfig

func init() {
	//Init config
	if err := env.Parse(&appConfig); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", appConfig)

	// Init Logrus, default to INFO
	if appConfig.Production {
		log.SetFormatter(&log.JSONFormatter{})

	} else {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05.00000",
		})

	}
	// log.SetFormatter(&log.JSONFormatter{})
	logLvl, err := log.ParseLevel(appConfig.LogLevel)
	if err != nil {
		// logLvl = log.InfoLevel
		logLvl = log.DebugLevel
	}
	log.SetLevel(logLvl)
	// log.SetReportCaller(true)
}

func main() {
	srv := fruitninja.FruitninjaSetup(&appConfig)
	srv.Logger.Fatal(srv.Start(":8080"))
}
