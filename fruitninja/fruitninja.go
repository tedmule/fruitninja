package fruitninja

import (
	"github.com/labstack/echo/v4"
)

func FruitninjaSetup() *echo.Echo {
	e := echo.New()

	// e.Use(middleware.Logger())

	e.GET("/", getFruitHandler)
	e.GET("/plenty", getPlentyOfFruitHandler)
	e.GET("/blade/:fruits", getBladeHandler)

	return e
}
