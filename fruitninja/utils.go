package fruitninja

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

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
		log.Error(err.Error())
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

func generatePetName() string {
	return petname.Generate(fruitNinjaSettings.Length, "_")
}

func produceFruit(fruitMap map[string]string, randomFruit ...bool) (fruit string) {
	var isRandom bool

	if len(randomFruit) > 0 {
		isRandom = randomFruit[0]
	}

	// If "isRandom" is true, generate random fruit
	if isRandom {
		fruits := make([]string, 0, len(fruitMap))
		for _, value := range fruitMap {
			fruits = append(fruits, value)
		}
		// Generate random fruit
		source := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(source)
		fruit = fruits[rnd.Intn(len(fruits))]

	} else {
		fruit = fruitMap[fruitNinjaSettings.Name]
	}
	return
}
