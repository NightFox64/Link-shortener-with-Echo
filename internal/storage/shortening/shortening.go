package shortening

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host     = "localhost"
	User     = "postgres"
	Password = "*******"
	Name     = "payment"
	Port     = "5432"
)

func Setup() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslname=disable",
		Host,
		Port,
		User,
		Name,
		Password)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
