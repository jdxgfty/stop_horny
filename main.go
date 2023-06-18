package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ALLOWED_HEADERS = []string{"GET", "HEAD"}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Match(ALLOWED_HEADERS, "/*.", func(c echo.Context) error {
		return c.Inline("stop_horny.jpg", "stop_horny.jpg")
	})

	log.Println("Starting server")

	e.Logger.Fatal(e.Start(":8080"))
}
