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
		db.AutoMigrate(&model.AllURLModel{})
		fmt.Println("Migrated")
	})
}

func TestCreateTabel(t *testing.T) {
	t.Run("Create urlTabel in DB", func(t *testing.T) {
		//connect to postgres
		db, err := shortening.Setup()
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Println("Connected")

		//create a urlTabel
		urlTabel := model.AllURLModel{
			OriginalURL: "qwertyuiop",
			ShortURL:    "qwe",
		}

		result, err := shortening.CreateURLTable(db, urlTabel)
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Println("urlTabel created", result)
	})
}

// just check it works in logic. I'm lazy to check it on existing tabel
func TestFindURL(t *testing.T) {
	t.Run("Find tabel with certain URL", func(t *testing.T) {
		//connect to postres
		db, err := shortening.Setup()
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Println("Connected")

		//select an url
		orig := "qwertyuiop"

		tabel, _ := shortening.FindShortURLWithOrig(db, orig)
		fmt.Println("Here's your tabel:", tabel)
	})
}

func TestUpdateURL(t *testing.T) {
	t.Run("Update tabel with certain URL", func(t *testing.T) {
		//connect to postres
		db, err := shortening.Setup()
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Println("Connected")

		//select an url
		orig := "qwertyuiop"
		urlTabel, _ := shortening.FindShortURLWithOrig(db, orig)
		fmt.Println("Your url is:", urlTabel)

		//update url with smth new
		updatedTabel, _ := shortening.UpdateShortURL(db, orig, model.AllURLModel{
			ShortURL: "123",
		})

		fmt.Println("Your url status now:", updatedTabel)
	})
}

func TestDeleteURL(t *testing.T) {
	t.Run("Update tabel with certain URL", func(t *testing.T) {
		//connect to postres
		db, err := shortening.Setup()
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Println("Connected")

		//select an url
		orig := "qwertyuiop"
		urlTabel, _ := shortening.FindShortURLWithOrig(db, orig)
		fmt.Println("Your url is:", urlTabel)

		//delete url
		shortening.DeleteURL(db, orig)
		fmt.Println("Your URL is deleted")
	})
}
