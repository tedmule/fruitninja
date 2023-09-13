package main

import (
	"github.com/daddvted/fruitninja/fruitninja"
)

func main() {
	srv := fruitninja.FruitninjaSetup()
	srv.Logger.Fatal(srv.Start(":8080"))
}
