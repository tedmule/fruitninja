package fruitninja

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"k8s.io/client-go/rest"
)

var config *rest.Config = &rest.Config{
	// Host:        "https://10.0.0.112:42883",
	// Host:        os.Getenv("K8S_API_URL"),
	// Host:        "https://kubernetes.default.svc",
	BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6Il8wcnFGQVcyQUs0bXdGZkhvT2ZhN01vTnRMazBnMHpDaHlHaHJBcEtqbXMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZXYiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoiZGV2LWNvbnRhaW5lci1zZWNyZXQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGV2LWNvbnRhaW5lci1zYSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjQyZmExZjBmLTVhZWUtNGYxMy05NDcxLWU1YTNlNTMzNDRmMiIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZXY6ZGV2LWNvbnRhaW5lci1zYSJ9.Cv5WFQJex8nMwbs7enlcZaDsk5bIe-dHmbbsxSJVnEiyA3PAwEI_ZW9Ry87rFBLRaJAvva75Hjrel734FBZhOuBvdMlaV1SVxQG6Ll_TDY_65q0D9NS578qT9cbYqpHlVYthzZ6iK2CEsE6-9t2zzjngYGJNLeIVXfTr3tXBFHc32bx48rA-RQBF9owUfc0rfoSGLOwwvyQ__hIxhUDJQSb63xQazQyox8zpAPjBXQXuTBD_igXlak932hN2MZsDK9WdCgU-Jvf8N4xEfSrcCFClwO8N8yQ7fvpuyuhH7II6Z7TBnlyYoZub7OOjq51Kt7nkSHxvcirU2QM7zEi_mw",
	TLSClientConfig: rest.TLSClientConfig{
		Insecure: true,
	},
}

func init() {
	k8sAPI := os.Getenv("K8S_API_SERVER")
	if strings.TrimSpace(k8sAPI) == "" {
		k8sAPI = "https://kubernetes.default.svc"
	}

	config.Host = k8sAPI

	fmt.Printf("%+v\n", config)
}

func FruitninjaSetup() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", getFruitHandler)
	e.GET("/plenty", getPlentyOfFruitHandler)
	e.GET("/blade/:fruits", getBladeHandler)

	return e
}
