package fruitninja

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	petname "github.com/dustinkirkland/golang-petname"
	log "github.com/sirupsen/logrus"
)

var fruitMap = map[string]string{
	"apple":      "🍎",
	"banana":     "🍌",
	"cherry":     "🍒",
	"coconut":    "🥥",
	"grape":      "🍇",
	"kiwi":       "🥝",
	"lemon":      "🍋",
	"mango":      "🥭",
	"orange":     "🍊",
	"peach":      "🍑",
	"pear":       "🍐",
	"pineapple":  "🍍",
	"strawberry": "🍓",
	"tomato":     "🍅",
	"watermelon": "🍉",
	"blade":      "🔪",
	"default":    "🐞",
}

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

func generateJabber() string {
	name := petname.Generate(fruitNinjaConfig.Length, "_")
	msg := strings.Repeat(fruitMap[fruitNinjaConfig.Name], fruitNinjaConfig.Count)

	return fmt.Sprintf("%s: %s", name, msg)
}
