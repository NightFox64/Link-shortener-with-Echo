package server

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/NightFox64/Link-shortener-with-Echo/internal/model"
	"github.com/NightFox64/Link-shortener-with-Echo/internal/shorten"
	"github.com/labstack/echo/v4"
	. "github.com/samber/mo"
)

type shortener interface {
	Shorten(context.Context, model.InputURLModel) (*model.AllURLModel, error)
}

type shortenRequest struct {
	URL        string `json:"url" validate:"required,url"`
	Identifier string `json:"identifier,omitempty" validate:"omitempty,alphanum"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url,omitempty"`
	Message  string `json:"message,omitempty"`
}

func HandlerShorten(shortener shortener) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req shortenRequest
		if err := c.Bind(&req); err != nil {
			return err
		}

		if err := c.Validate(req); err != nil {
			return err
		}

		identifier := None[string]()
		if strings.TrimSpace(req.Identifier) != "" {
			identifier = Some(req.Identifier)
		}

		input := model.InputURLModel{
			RawURL:     req.URL,
			Identifier: identifier,
		}

		shortening, err := shortener.Shorten(c.Request().Context(), input)
		if err != nil {
			return err
		}

		shortURL, err := shorten.PrependBaseURL(shortening.OriginalURL, shortening.Identifier)
		if err != nil {
			log.Println("shorten.PrependBaseURL:", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(
			http.StatusOK,
			shortenResponse{ShortURL: shortURL},
		)
	}
}
