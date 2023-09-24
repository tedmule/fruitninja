package fruitninja

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
)

type FruitNinjaConfig struct {
	Name       string `env:"FRUIT_NINJA_NAME" envDefault:"default"`
	Count      int    `env:"FRUIT_NINJA_COUNT" envDefault:"1"`
	LogLevel   string `env:"FRUIT_NINJA_LOG_LEVEL" envDefault:"debug"`
	K8SAPI     string `env:"FRUIT_NINJA_K8A_API" envDefault:"https://kubernetes.default.svc"`
	Production bool   `env:"FRUIT_NINJA_PRODUCTION" envDefault:"false"`
}

var k8sConfig *rest.Config

var fruitNinjaConfig FruitNinjaConfig

func FruitninjaSetup(config *FruitNinjaConfig) *echo.Echo {
	fruitNinjaConfig = *config

	e := echo.New()

	// e.Use(middleware.Logger())
	// log.SetFormatter(&log.JSONFormatter{})
	// log.SetReportCaller(true)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogUserAgent: true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(log.Fields{
				"URI":       values.URI,
				"status":    values.Status,
				"latency":   values.Latency,
				"reqid":     values.RequestID,
				"useragent": values.UserAgent,
			}).Info("HTTP")

			return nil
		},
	}))

	e.GET("/", getFruitHandler)
	e.GET("/plenty", getPlentyOfFruitHandler)
	e.GET("/blade/:fruits", getBladeHandler)

	return e
}
