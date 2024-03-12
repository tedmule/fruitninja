package main

import (
	"github.com/caarlos0/env/v9"
	"github.com/daddvted/fruitninja/data"
	"github.com/daddvted/fruitninja/fruitninja"
	log "github.com/sirupsen/logrus"
)

var settings fruitninja.FruitNinjaSettings

func init() {
	//Init config
	if err := env.Parse(&settings); err != nil {
		log.Error(err.Error())
	}
	log.Debugf("%+v\n", settings)

	// Init Logrus, default to INFO
	if settings.Production {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05.00000",
		})

	}
	// log.SetFormatter(&log.JSONFormatter{})
	logLvl, err := log.ParseLevel(settings.LogLevel)
	if err != nil {
		// logLvl = log.InfoLevel
		logLvl = log.DebugLevel
	}
	log.SetLevel(logLvl)
	log.SetReportCaller(true)
}

func main() {
	// Connect to Redis at start
	cache, err := data.NewRedisClient(settings.RedisAddr, settings.RedisDB)
	if err != nil {
		log.Errorf("Failed to connect to Redis: %s", err.Error())
	}

	// Connect to MySQL at start
	mysql, err := data.NewMysqlClient(settings.MySQLHost, settings.MySQLUsername, settings.MySQLPassword, settings.MySQLDB)
	if err != nil {
		log.Errorf("Failed to connect to MySQL: %s", err.Error())
	}

	httpSrv := fruitninja.FruitNinjaSetup(&settings, cache, mysql)
	log.Infof("Fruitninja runs in %s mode.", settings.Mode)
	httpSrv.Logger.Fatal(httpSrv.Start(":8080"))
}
