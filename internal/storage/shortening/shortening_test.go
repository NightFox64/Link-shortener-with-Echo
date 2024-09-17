package shortening_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/NightFox64/Link-shortener-with-Echo/internal/storage/shortening"
)

func TestConnection(t *testing.T) {
	t.Run("connects to DB", func(t *testing.T) {
		_, err := shortening.Setup()
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Println("Connect success")
	})
}
