package main

import (
	"github.com/NightFox64/Link-shortener-with-Echo/internal/server"
	"github.com/NightFox64/Link-shortener-with-Echo/internal/storage/shortening"
	"github.com/labstack/echo/v4"
)

// var DataBase *gorm.DB
var DataBase, _ = shortening.Setup()

func main() {
	//url := shorten.Shorten(shorten.GenerateNewURL())

	e := echo.New()

	e.POST("/shorten", server.ShortenURLHandler)
	e.GET("/:shortened", server.RedirectURLHandler)

	e.Start(":8080")
}
