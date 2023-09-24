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
	return c.String(http.StatusOK, fmt.Sprintf("%s\n", fruitMap[fruitNinjaConfig.Name]))
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

	return c.String(http.StatusOK, strings.Join(skewer, "->"))
}
