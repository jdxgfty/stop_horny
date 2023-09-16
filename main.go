package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"embed"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

var ALLOWED_HEADERS = []string{"GET", "HEAD"}

//go:embed stop_horny.jpg
var imgFS embed.FS

func main() {
	e := echo.New()

	if slices.Contains(os.Args, "-l") {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	e.Match(ALLOWED_HEADERS, "/*.", func(c echo.Context) error {
		return inline(c, "stop_horny.jpg", "stop_horny.jpg")
	})

	log.Println("Starting server")

	e.Logger.Error(e.Start(":8080"))
}

func inline(c echo.Context, file, name string) error {
	return contentDisposition(c, file, name, "inline")
}

func contentDisposition(c echo.Context, file, name, dispositionType string) error {
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("%s; filename=%q", dispositionType, name))
	return c.FileFS(name, imgFS)
}
