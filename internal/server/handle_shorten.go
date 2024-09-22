package server

import (
	"fmt"
	"net/http"

	"github.com/NightFox64/Link-shortener-with-Echo/internal/model"
	"github.com/NightFox64/Link-shortener-with-Echo/internal/shorten"
	"github.com/NightFox64/Link-shortener-with-Echo/internal/storage/shortening"
	"github.com/labstack/echo/v4"
)

type longURL struct {
	Url string `json:"url"`
}

func ShortenURLHandler(c echo.Context) error {
	var url longURL
	if err := c.Bind(&url); err != nil {
		return err
	}
	fmt.Println(url.Url)

	shortened := shorten.Shorten(url.Url)

	urlTabel := model.AllURLModel{
		OriginalURL: url.Url,
		ShortURL:    shortened,
	}

	_, err0 := shortening.CreateURLTable(urlTabel)
	if err0 != nil {
		return err0
	}

	return c.JSON(http.StatusOK, urlTabel)
}

func RedirectURLHandler(c echo.Context) error {
	data := c.Param("shortened")
	urlTabel, err := shortening.FindOrigURLWithShort(data)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, urlTabel.OriginalURL)
}
