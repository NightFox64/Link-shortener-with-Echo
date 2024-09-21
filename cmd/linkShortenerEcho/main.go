package main

import (
	"github.com/NightFox64/Link-shortener-with-Echo/internal/server"
	"github.com/NightFox64/Link-shortener-with-Echo/internal/shorten"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	url := shorten.Shorten(shorten.GenerateNewURL())

	e := echo.New()

	e.POST("/shorten", server.ShortenURLHandler)
	e.Pre(middleware.HTTPSRedirect())

	e.Logger.Fatal(e.Start(":8080"))
}
