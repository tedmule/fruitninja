package fruitninja

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/daddvted/fruitninja/data"
	"github.com/labstack/echo/v4"
	"github.com/mileusna/useragent"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getK8sFruitHandler(c echo.Context) error {
	log.Infof("Request for [%s] service\n", fruitNinjaSettings.Name)
	url := c.Request().URL.Path
	log.Debugf("Request URL: %s\n", url)

	if strings.TrimSpace(url) == "/" {
		msg := strings.Repeat(fruitMap[fruitNinjaSettings.Name], fruitNinjaSettings.Count)
		return c.String(http.StatusOK, fmt.Sprintf("%s\n", msg))
	}

	splitedURL := strings.SplitN(strings.Trim(url, "/"), "/", 2)
	serviceLength := len(splitedURL)
	log.Debugf("Splited URL length: %d\n", serviceLength)

	skewer := []string{}
	// Append itself to skewer
	skewer = append(skewer, fruitMap[fruitNinjaSettings.Name])
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

func getK8sBladeHandler(c echo.Context) error {
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

func getJabberHandler(c echo.Context) error {
	var jabberText string
	var cacheText string
	var dbText string
	var respText string

	fmt.Println(c.Request().UserAgent())
	ua := useragent.Parse(c.Request().UserAgent())

	fruitName := produceFruit(fruitMap, true)

	// Cache
	if fruitNinjaCache == nil {
		// Try to connect Redis again.
		redis, err := data.NewRedisClient(fruitNinjaSettings.RedisAddr, fruitNinjaSettings.RedisDB)
		if err != nil {
			log.Errorf("Failed to connect to Redis: %s", err.Error())
			cacheText = "Failed to connect to redis"
		} else {
			fruitNinjaCache = redis
			fruitNinjaCache.AppendKey("fruits", fruitName)
			cacheText = fruitNinjaCache.GetKey("fruits")
		}
	} else {
		fruitNinjaCache.AppendKey("fruits", fruitName)
		cacheText = fruitNinjaCache.GetKey("fruits")
	}
	log.Debug(fruitNinjaMysql)

	// DB
	if fruitNinjaMysql != nil {
		_, err := fruitNinjaMysql.GetSingleFruit(fruitName)
		if err != nil {
			log.Error(err)
			_, err := fruitNinjaMysql.AddFruit(fruitName)
			if err != nil {
				log.Error(err)
			}
		} else {
			err := fruitNinjaMysql.AddAmount(fruitName)
			if err != nil {
				log.Error(err)
			}
		}

		// Query all fruits
		fruits, err := fruitNinjaMysql.GetFruits()
		if err != nil {
			log.Error(err)
			dbText = err.Error()
		} else {
			var text []string
			for _, fruit := range fruits {
				text = append(text, fmt.Sprintf("%s: %d", fruit.Name, fruit.Amount))
			}
			dbText = strings.Join(text, "\n")
		}

	} else {
		// re-connect
		mysql, err := data.NewMysqlClient(fruitNinjaSettings.MySQLHost, fruitNinjaSettings.MySQLUsername, fruitNinjaSettings.MySQLPassword, fruitNinjaSettings.MySQLDB)
		if err != nil {
			log.Errorf("Failed to connect to MySQL again: %s", err.Error())
			dbText = "Failed to connect to MySQL"
		}
		fruitNinjaMysql = mysql
	}

	jabberText = fmt.Sprintf("%s: %s", generatePetName(true), strings.Repeat(fruitName, fruitNinjaSettings.Count))

	if ua.Name == "curl" {
		version := fmt.Sprintf("--- Version: %s ---", Version)
		respText = fmt.Sprintf("%s\n%s\nCACHE:%s\nDB:\n%s\n", version, jabberText, cacheText, dbText)
		return c.String(http.StatusOK, respText)
	} else {
		version := fmt.Sprintf("<h1>Version: %s</h1>", Version)
		respText = fmt.Sprintf("%s\n%s<br/>CACHE:%s<br/>DB:<br/>%s<br/>", version, jabberText, cacheText, dbText)
		return c.HTML(http.StatusOK, respText)
	}
}

func wsHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		msg := "Welcom to FruitNinja"
		for {
			// Write
			err := websocket.Message.Send(ws, msg)
			if err != nil {
				log.Error(err)
			}

			// Read
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				log.Error(err)
			}
			fmt.Printf("received %s from client\n", msg)
			msg = fruitMap[msg]
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
