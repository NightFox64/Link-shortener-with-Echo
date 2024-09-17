package model

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("not found")
	ErrIDExist  = errors.New("this id already exists")
)

type ShortenedURLModel struct {
	gorm.Model
	originalURL string
	shortURL    string `gorm: unique index`
}
