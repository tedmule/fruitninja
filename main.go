package main

import (
	"github.com/daddvted/fruitninja/fruitninja"
)

func main() {
	srv := fruitninja.EchoSetup()
	srv.Logger.Fatal(srv.Start(":8080"))
}
