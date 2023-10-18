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
	"blade":      "🔪",
	"default":    "🐞",
}

func getFruitHandler(c echo.Context) error {
	log.Infof("Request for [%s] service\n", fruitNinjaConfig.Name)
	url := c.Request().URL.Path
	log.Debugf("Request URL: %s\n", url)

	if strings.TrimSpace(url) == "/" {
		msg := strings.Repeat(fruitMap[fruitNinjaConfig.Name], fruitNinjaConfig.Count)
		return c.String(http.StatusOK, fmt.Sprintf("%s\n", msg))
	}

	splitedURL := strings.SplitN(strings.Trim(url, "/"), "/", 2)
	serviceLength := len(splitedURL)
	log.Debugf("Splited URL length: %d\n", serviceLength)

	skewer := []string{}
	// Append itself to skewer
	skewer = append(skewer, fruitMap[fruitNinjaConfig.Name])
	ns := getNamespace()

	var urlRemainder, serviceURL string

	for serviceLength > 1 {
		nextService := splitedURL[0]
		if serviceLength > 1 {
			urlRemainder = splitedURL[1]
			serviceURL = "http://" + nextService + "." + ns + ".svc.cluster.local/" + urlRemainder
		} else {
			// If serviceLength == 1, no need to append urlRemainder
			serviceURL = "http://" + nextService + "." + ns + ".svc.cluster.local/"
		}
		log.Debugf("Next service: %s\n", nextService)
		log.Debugf("URL remainder: %s\n", urlRemainder)
		log.Debugf("Next service url: %s\n", serviceURL)

		fruitEmoji, ok := getServingFruit(serviceURL)
		if ok {
			skewer = append(skewer, strings.TrimSpace(fruitEmoji))
			break
		} else {
			// Enclose fruit with square bracket,
			// when fruit emoji not return successfully.
			skewer = append(skewer, fmt.Sprintf("[%s]", nextService))
			splitedURL = strings.SplitN(urlRemainder, "/", 2)
			fmt.Printf("splited url: %+v\n", splitedURL)
			serviceLength = len(splitedURL)
			fmt.Printf("service LENGTH: %d\n", serviceLength)
		}

	}

	if serviceLength == 1 {
		nextService := splitedURL[0]
		serviceURL := "http://" + nextService + "." + ns + ".svc.cluster.local/"
		fmt.Printf("next service url: %s\n", serviceURL)
		fruitEmoji, ok := getServingFruit(serviceURL)
		if ok {
			// skewer = append(skewer, strings.TrimSpace(fruitEmoji))
			fmt.Println(fruitEmoji)
			skewer = append(skewer, strings.TrimSpace(fruitEmoji))
		} else {
			// Enclose fruit with square bracket,
			// when fruit emoji not return successfully.
			skewer = append(skewer, fmt.Sprintf("[%s]", nextService))
		}

	}
	bladeString := strings.Join(skewer, "->")
	// return c.String(http.StatusOK, strings.Join(skewer, "->"))
	return c.String(http.StatusOK, fmt.Sprintf("%s\n", bladeString))

}

func getJabberHandler(c echo.Context) error {
	name := petname.Generate(3, "_")
	msg := strings.Repeat(fruitMap[fruitNinjaConfig.Name], fruitNinjaConfig.Count)

	return c.String(http.StatusOK, fmt.Sprintf("%s: %s\n", name, msg))
}

func getBladeHandler(c echo.Context) error {
	url := c.Request().URL.Path
	log.Debugf("Request URL: %s\n", url)
	queryStr := c.Param("fruits")
	log.Debugf("Query string: %s\n", queryStr)

	splitedURL := strings.SplitN(strings.Trim(queryStr, "/"), "/", 2)
	log.Debugf("Splited URL: %q\n", splitedURL)
	fruitNO := len(splitedURL)
	log.Debugf("No of fruits: %d\n", fruitNO)
	ns := getNamespace()

	var serviceURL string

	if fruitNO > 1 {
		serviceURL = "http://" + splitedURL[0] + "." + ns + ".svc.cluster.local/" + splitedURL[1]
	} else {
		// If serviceLength == 1, no need to append urlRemainder
		serviceURL = "http://" + splitedURL[0] + "." + ns + ".svc.cluster.local/"
	}
	fruitEmoji, ok := getServingFruit(serviceURL)
	if ok {
		return c.String(http.StatusOK, fruitEmoji)
	} else {
		return c.String(400, fmt.Sprintf("%s\n", fruitMap["blade"]))
	}
}
