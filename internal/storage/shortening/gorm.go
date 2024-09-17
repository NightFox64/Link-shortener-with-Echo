package shortening

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDB() {
	dsn := "host=localhost user=user dbname=db password=password sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&ShortenedURLModel{})
}
