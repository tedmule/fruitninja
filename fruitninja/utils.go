package fruitninja

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func getMatchedService(name string, services *[]string) (string, bool) {
	for _, service := range *services {
		if strings.Contains(service, name) {
			return service, true
		}
	}
	return "", false
}

func getServingFruit(url string) (string, bool) {
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Printf("%+v\n", err)
		log.Error(err.Error())
		return "", false
	}

	if resp.StatusCode != 200 {
		log.Info(resp.StatusCode)
		return "", false
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	return string(body), true
}

func getNamespace() string {
	namespace, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		log.Error(err.Error())
		// Return default namespace when encounting error
		return "default"
	}
	return string(namespace)
}

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return "NO_HOSTNAME"
	} else {
		return name
	}

}
