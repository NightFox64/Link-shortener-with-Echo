package shortening

import (
	"errors"
	"fmt"

	"github.com/NightFox64/Link-shortener-with-Echo/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host     = "localhost"
	User     = "postgres"
	Password = "Ichiho64"
	Name     = "postgres"
	Port     = "5432"
)

func Setup() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		Host,
		Port,
		User,
		Name,
		Password)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.AllURLModel{})

	return db, nil

}

func CreateURLTable(urlTabel model.AllURLModel) (int64, error) {
	result := GlobalDB.Create(&urlTabel)
	if result.RowsAffected == 0 {
		return 0, errors.New("urlTabel not created")
	}
	return result.RowsAffected, nil
}

// Idk what I did
func FindShortURLWithOrig(orig string) (model.AllURLModel, error) {
	var urlTabel model.AllURLModel
	result := GlobalDB.First(&urlTabel, "original_url = ?", orig)
	if result.RowsAffected == 0 {
		return model.AllURLModel{}, errors.New("data not found")
	}
	return urlTabel, nil
}

func FindOrigURLWithShort(short string) (model.AllURLModel, error) {
	var urlTabel model.AllURLModel
	result := GlobalDB.First(&urlTabel, "short_url = ?", short)
	if result.RowsAffected == 0 {
		return model.AllURLModel{}, errors.New("data not found")
	}
	return urlTabel, nil
}

// you send db, orig url and tabel with changed stats
func UpdateShortURL(orig string, urlTabel model.AllURLModel) (model.AllURLModel, error) {
	var updateURLTabel model.AllURLModel
	result := GlobalDB.Model(&updateURLTabel).Where("original_url = ?", orig).Updates(urlTabel)
	if result.RowsAffected == 0 {
		return model.AllURLModel{}, errors.New("can't update")
	}
	return updateURLTabel, nil
}

func DeleteURL(orig string) (int64, error) {
	var deletedTabel model.AllURLModel
	result := GlobalDB.Where("original_url = ?", orig).Delete(&deletedTabel)
	if result.RowsAffected == 0 {
		return 0, errors.New("https://youtu.be/nwuW98yLsgY?feature=shared")
	}
	return result.RowsAffected, nil
}
