package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var ALLOWED_HEADERS = []string{"GET", "HEAD"}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	r.Match(ALLOWED_HEADERS, "/*.", func(c *gin.Context) {
		c.File("stop_horny.jpg")
	})

	log.Println("Starting server")

	r.Run(":8080")
}
