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

func EchoSetup() *echo.Echo {
	e := echo.New()

	// r.GET("/:fruit/:count", func(ctx *gin.Context) {
	e.GET("/", func(c echo.Context) error {
		var msg string

		fruit := os.Getenv("FRUIT_NINJA_NAME")
		count := os.Getenv("FRUIT_NINJA_COUNT")
		name := petname.Generate(3, "_")

		// fruit := ctx.Param("fruit")
		cnt, err := strconv.Atoi(count)
		if err != nil {
			fmt.Printf("%s: %s\n", "🐞", err.Error())
			cnt = 1
		}

		switch fruit {
		case "apple":
			msg = strings.Repeat("🍎", cnt)
		case "banana":
			msg = strings.Repeat("🍌", cnt)
		case "orange":
			msg = strings.Repeat("🍊", cnt)
		case "watermelon":
			msg = strings.Repeat("🍉", cnt)
		case "pear":
			msg = strings.Repeat("🍐", cnt)
		case "cherry":
			msg = strings.Repeat("🍒", cnt)
		case "strawberry":
			msg = strings.Repeat("🍓", cnt)
		case "kiwi":
			msg = strings.Repeat("🥝", cnt)
		default:
			msg = "🐞"
		}

		return c.String(http.StatusOK, fmt.Sprintf("%s: %s\n", name, msg))
	})

	return e
}
