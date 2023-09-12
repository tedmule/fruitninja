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

func EchoSetup() *echo.Echo {
	e := echo.New()

	// e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {

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
	})

	return e
}
