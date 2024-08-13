package fruitninja

import (
	"github.com/daddvted/fruitninja/data"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
)

type FruitNinjaSettings struct {
	// Context    string `env:"FRUIT_NINJA_CONTEXT" envDefault:""`
	Mode          string `env:"FRUIT_NINJA_MODE" envDefault:"default"`
	Length        int    `env:"FRUIT_NINJA_JABBER_WORD" envDefault:"2"`
	Name          string `env:"FRUIT_NINJA_NAME" envDefault:"default"`
	Count         int    `env:"FRUIT_NINJA_COUNT" envDefault:"1"`
	LogLevel      string `env:"FRUIT_NINJA_LOG_LEVEL" envDefault:"debug"`
	K8SAPI        string `env:"FRUIT_NINJA_K8A_API" envDefault:"https://kubernetes.default.svc"`
	Production    bool   `env:"FRUIT_NINJA_PROD" envDefault:"false"`
	K8SToken      string `env:"FRUIT_NINJA_K8S_TOKEN" envDefault:"eyJhbGciOiJSUzI1NiIsImtpZCI6InRBb1JyNzRaa3VYZmV6cmk4bHZybGJZcjVpOGN4cDhCSEtCdEJQMnp1RWMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZXYiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoiZGV2LWNvbnRhaW5lci1zZWNyZXQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGV2LWNvbnRhaW5lci1zYSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjExMWY0OGRmLTFkYWEtNDljOS1hMzIzLTI0Nzc3ZWE0Y2U0ZCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZXY6ZGV2LWNvbnRhaW5lci1zYSJ9.o0ZKu_ziOO3-GJ_kzYDnNq3UslhjRkue0TJWFAC9wgAgndhQi37r6-HwtMx3syHnC8Q5sNdG_Df0vYAKSH5PjgA2RqbMIoOWUwRxEDIwNBHHZ9xJrOu4gCZoxWqHgBskmjsqE5zVw5D6ksltAEZKFke15t2NlYuiiaz1Mj9mcEdUk7ryo5Z18VGKe6lsdbqfu_6GkUvN5NvzvoZcSrnc6VTGxuBV_c1Mfhk0lJpIlzEZjjDCpi6w-V3aH1oIJE5xmBxSOo9i8GRCV1SmEMsOErF9Qsc2QRwIiuIe4R4ALS-xSxqbrDBEAnI95feZDlsJU8yrqMsm0zxpkpHWSHQ13Q"`
	RedisAddr     string `env:"FRUIT_NINJA_REDIS_ADDR" envDefault:"localhost:6379"`
	RedisPassword string `env:"FRUIT_NINJA_REDIS_PSD" envDefault:""`
	RedisDB       int    `env:"FRUIT_NINJA_REDIS_DB" envDefault:"0"`
	MySQLHost     string `env:"FRUIT_NINJA_MYSQL_HOST" envDefault:"localhost:3306"`
	MySQLUsername string `env:"FRUIT_NINJA_MYSQL_USERNAME" envDefault:"root"`
	MySQLPassword string `env:"FRUIT_NINJA_MYSQL_PASSWORD" envDefault:"root"`
	MySQLDB       string `env:"FRUIT_NINJA_MYSQL_DB" envDefault:"fruit"`
}

var (
	k8sConfig *rest.Config

	fruitNinjaSettings *FruitNinjaSettings
	fruitNinjaCache    *data.Cache
	fruitNinjaMysql    *data.DB
)

func FruitNinjaSetup(settings *FruitNinjaSettings, cache *data.Cache, mysql *data.DB) *echo.Echo {
	e := echo.New()
	fruitNinjaSettings = settings
	fruitNinjaCache = cache
	fruitNinjaMysql = mysql

	// e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.IPExtractor = echo.ExtractIPFromRealIPHeader()

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

	// "k8s" mode is used for service mesh demo
	if settings.Mode == "k8s" {
		e.GET("/*", getK8sFruitHandler)
		e.GET("/blade/:fruits", getK8sBladeHandler)

	} else {
		e.File("/", "static/html/index.html")
		e.GET("/jabber", getJabberHandler)
		e.GET("/hello", getHelloHandler)
		e.GET("/fruit", getFruitHandler)
		e.GET("/ws", wsHandler)
	}

	return e
}
