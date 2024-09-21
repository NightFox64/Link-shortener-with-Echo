package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/shorten", shortenURL)
	e.GET("/:shortened", redirectURL)

	e.Logger.Fatal(e.Start(":8080"))
}
