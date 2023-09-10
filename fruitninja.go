package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	r := gin.Default()

	// r.GET("/:fruit/:count", func(ctx *gin.Context) {
	r.GET("/", func(ctx *gin.Context) {
		var msg string

		fruit := os.Getenv("FRUIT_NINJA_NAME")
		fmt.Println(fruit)
		count := os.Getenv("FRUIT_NINJA_COUNT")

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

		ctx.String(200, msg)
	})

	return r
}
