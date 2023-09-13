package fruitninja

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getMatchedService(name string, services *[]string) (string, bool) {
	fmt.Println(services)

	for _, service := range *services {
		if strings.Contains(service, name) {
			return service, true
		}
	}
	return "", false
}

func getServingFruit(url string) (string, bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return "", false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	return string(body), true
}
