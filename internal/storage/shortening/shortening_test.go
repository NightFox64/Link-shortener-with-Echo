package shortening_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/NightFox64/Link-shortener-with-Echo/internal/model"
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

func TestMigrate(t *testing.T) {
	t.Run("Migrate DB with model", func(t *testing.T) {
		db, err := shortening.Setup()
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Println("Connected")
		db.AutoMigrate(model.AllURLModel{})
		fmt.Println("Migrated")
	})
}
