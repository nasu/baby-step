package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", index())
	log.Printf("Listening on port %d", 8080)
	e.Start(":8080")
}

func index() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("Called index page.")
		return c.String(http.StatusOK, "Hello, GAE 2nd.")
	}
}
