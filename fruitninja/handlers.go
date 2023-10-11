package fruitninja

import (
	"fmt"
	"net/http"
	"strings"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/labstack/echo/v4"
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
	"default":    "🐞",
}

func getFruitHandler(c echo.Context) error {
	log.Infof("Request for [%s] service\n", fruitNinjaConfig.Name)
	url := c.Request().URL.Path
	fmt.Println(url)
	// currentNS := getNamespace()
	// fmt.Printf("Current namespace: %s\n", currentNS)
	// fmt.Println(strings.SplitN(strings.Trim(url, "/"), "/", 2))
	splitedURL := strings.SplitN(strings.Trim(url, "/"), "/", 2)
	serviceLength := len(splitedURL)
	fmt.Printf("START LENGTH: %d\n", serviceLength)
	skewer := []string{}
	skewer = append(skewer, fruitMap[fruitNinjaConfig.Name])
	ns := getNamespace()

	for serviceLength > 1 {
		urlRemainder := splitedURL[1]
		nextService := splitedURL[0]
		fmt.Printf("url remainder: %s\n", urlRemainder)
		fmt.Printf("Next service: %s\n", nextService)

		serviceURL := "http://" + nextService + "." + ns + ".svc.cluster.local/" + urlRemainder
		fmt.Printf("next service url: %s\n", serviceURL)

		fruitEmoji, ok := getServingFruit(serviceURL)
		fmt.Printf("OK: %t\n", ok)
		if ok {
			// skewer = append(skewer, strings.TrimSpace(fruitEmoji))
			fmt.Println(fruitEmoji)
			skewer = append(skewer, strings.TrimSpace(fruitEmoji))
			// return c.String(http.StatusOK, fmt.Sprintf("%s->%s\n", fruitMap[fruitNinjaConfig.Name], fruitEmoji))
		} else {
			// Enclose fruit with square bracket,
			// when fruit emoji not return successfully.
			skewer = append(skewer, fmt.Sprintf("[%s]", nextService))
			splitedURL = strings.SplitN(strings.Trim(urlRemainder, "/"), "/", 2)
			serviceLength = len(splitedURL)
			fmt.Printf("service LENGTH: %d\n", serviceLength)
		}

	}
	bladeString := strings.Join(skewer, "->")
	// return c.String(http.StatusOK, strings.Join(skewer, "->"))
	return c.String(http.StatusOK, fmt.Sprintf("%s\n", bladeString))

}

func getPlentyOfFruitHandler(c echo.Context) error {
	name := petname.Generate(3, "_")
	msg := strings.Repeat(fruitMap[fruitNinjaConfig.Name], fruitNinjaConfig.Count)

	return c.String(http.StatusOK, fmt.Sprintf("%s: %s\n", name, msg))
}

func getBladeHandler(c echo.Context) error {
	skewer := []string{}

	fruits := strings.Split(c.Param("fruits"), "/")
	services := getK8SService()
	log.Debug(services)

	for _, fruit := range fruits {
		matchedSvc, found := getMatchedService(fruit, &services)

		if found {
			log.Infof("Matched service for %s is %s\n", fruit, matchedSvc)
			url := "http://" + matchedSvc + ".fruitninja"
			fruitEmoji, ok := getServingFruit(url)
			if ok {
				skewer = append(skewer, strings.TrimSpace(fruitEmoji))
			} else {
				// Enclose fruit with square bracket,
				// when fruit emoji not return successfully.
				skewer = append(skewer, fmt.Sprintf("[%s]", fruit))
			}
		} else {
			skewer = append(skewer, fmt.Sprintf("[%s]", fruit))
		}
	}

	bladeString := strings.Join(skewer, "->")
	// return c.String(http.StatusOK, strings.Join(skewer, "->"))
	return c.String(http.StatusOK, fmt.Sprintf("%s\n", bladeString))
}
