package server

import (
	"net/http"

	"github.com/NightFox64/Link-shortener-with-Echo/internal/model"
	"github.com/NightFox64/Link-shortener-with-Echo/internal/shorten"
	"github.com/NightFox64/Link-shortener-with-Echo/internal/storage/shortening"
	"github.com/labstack/echo/v4"
)

func ShortenURLHandler(c echo.Context) error {
	//longURL := c.FormValue("url")

	longURL := ""
	if err := c.Bind(longURL); err != nil {
		return err
	}

	//if longURL == "" {
	//	return c.JSON(http.StatusBadRequest, "error: URL is required")
	//}

	shortened := shorten.Shorten(longURL)

	db, err := shortening.Setup()
	if err != nil {
		return err
	}

	urlTabel := model.AllURLModel{
		OriginalURL: longURL,
		ShortURL:    shortened,
	}

	_, err0 := shortening.CreateURLTable(db, urlTabel)
	if err0 != nil {
		return err
	}

	return c.JSON(http.StatusOK, urlTabel)
}

func RedirectURLHandler(c echo.Context) error {
	data := c.Param("shortened")
	return c.Redirect(http.StatusFound, data)
}
