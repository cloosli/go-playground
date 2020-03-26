package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015", // endpoint without http
	})
	if err != nil {
		log.Fatalln(err)
	}

	res, err := r.Expr("Hello World").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	r.TableCreate("test", r.TableCreateOpts{
		Shards:   2,
		Replicas: 3,
	})

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)

	// Output:
	// Hello World
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}",` +
			`"method":"${method}","uri":"${uri}",` +
			`"status":${status},"error":"${error}","latency_human":"${latency_human}"` +
			"\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	var fEnvironment string
	flag.StringVar(&fEnvironment, "env", "DEV", "environment: DEV (development), STG (staging), PRD (production)")
	flag.Parse()
	if fEnvironment == "DEV" {
		e.Logger.Fatal(e.Start("localhost:1323"))
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
