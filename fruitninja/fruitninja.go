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

var k8sConfig *rest.Config = &rest.Config{
	// Host:        fruitNinjaConfig.K8SAPI,
	BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6Il8wcnFGQVcyQUs0bXdGZkhvT2ZhN01vTnRMazBnMHpDaHlHaHJBcEtqbXMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZXYiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoiZGV2LWNvbnRhaW5lci1zZWNyZXQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGV2LWNvbnRhaW5lci1zYSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjQyZmExZjBmLTVhZWUtNGYxMy05NDcxLWU1YTNlNTMzNDRmMiIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZXY6ZGV2LWNvbnRhaW5lci1zYSJ9.Cv5WFQJex8nMwbs7enlcZaDsk5bIe-dHmbbsxSJVnEiyA3PAwEI_ZW9Ry87rFBLRaJAvva75Hjrel734FBZhOuBvdMlaV1SVxQG6Ll_TDY_65q0D9NS578qT9cbYqpHlVYthzZ6iK2CEsE6-9t2zzjngYGJNLeIVXfTr3tXBFHc32bx48rA-RQBF9owUfc0rfoSGLOwwvyQ__hIxhUDJQSb63xQazQyox8zpAPjBXQXuTBD_igXlak932hN2MZsDK9WdCgU-Jvf8N4xEfSrcCFClwO8N8yQ7fvpuyuhH7II6Z7TBnlyYoZub7OOjq51Kt7nkSHxvcirU2QM7zEi_mw",
	TLSClientConfig: rest.TLSClientConfig{
		Insecure: true,
	},
}
var fruitNinjaConfig FruitNinjaConfig

func FruitninjaSetup(config *FruitNinjaConfig) *echo.Echo {
	fruitNinjaConfig = *config
	k8sConfig.Host = fruitNinjaConfig.K8SAPI

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
