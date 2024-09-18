package model

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrIDExist  = errors.New("this id already exists")
)

type ShortenedURLModel struct {
	originalURL string
	shortURL    string `gorm:"unique_short_index"`
}
