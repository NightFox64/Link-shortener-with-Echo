package shorten_test

import (
	"fmt"
	"testing"

	"github.com/NightFox64/Link-shortener-with-Echo/internal/shorten"
)

func TestShorten(t *testing.T) {
	t.Run("returns a short URL", func(t *testing.T) {

		//let's see if it generates a new thing and shorten it
		url := shorten.GenerateNewURL()
		fmt.Println("Random URL:", url)

		shortenURL := shorten.Shorten(url)
		fmt.Println("Now it's shorten:", shortenURL)
	})
}
