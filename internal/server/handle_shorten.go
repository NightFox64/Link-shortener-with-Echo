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
	url string
}

func ShortenURLHandler(c echo.Context) error {
	url := new(longURL)
	//if err := c.Bind(&url); err != nil {
	//	return err
	//}

	err1 := (&echo.DefaultBinder{}).BindBody(c, &url)
	if err1 != nil {
		return err1
	}

	fmt.Println(url.url)

	shortened := shorten.Shorten(url.url)

	db, err := shortening.Setup()
	if err != nil {
		return err
	}

	urlTabel := model.AllURLModel{
		OriginalURL: url.url,
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
