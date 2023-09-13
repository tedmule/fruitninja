package fruitninja

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/labstack/echo/v4"
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
	fruit := os.Getenv("FRUIT_NINJA_NAME")
	if strings.TrimSpace(fruit) == "" {
		fruit = "default"
	}
	return c.String(http.StatusOK, fmt.Sprintf("%s\n", fruitMap[fruit]))

}

func getPlentyOfFruitHandler(c echo.Context) error {

	fruit := os.Getenv("FRUIT_NINJA_NAME")
	if strings.TrimSpace(fruit) == "" {
		fruit = "default"
	}
	count := os.Getenv("FRUIT_NINJA_COUNT")
	name := petname.Generate(3, "_")

	cnt, err := strconv.Atoi(count)
	if err != nil {
		fmt.Printf("%s: %s\n", "🐞", err.Error())
		cnt = 1
	}

	msg := strings.Repeat(fruitMap[fruit], cnt)

	return c.String(http.StatusOK, fmt.Sprintf("%s: %s\n", name, msg))
}

func getBladeHandler(c echo.Context) error {
	skewer := []string{}

	fruits := strings.Split(c.Param("fruits"), "/")
	// services := getK8SService("fruitninja")

	// for _, fruit := range fruits {
	// 	matched, found := getMatchedService(fruit, &services)

	// 	if found {
	// 		fmt.Printf("Matched service for %s is %s\n", fruit, matched)
	// 	} else {
	// 		fmt.Printf("Matched service for %s is not found\n", fruit)
	// 	}
	// }
	fruitEmoji, ok := getServingFruit("http://localhost:8080/")

	return c.String(http.StatusOK, strings.Join(fruits, "->"))
}
