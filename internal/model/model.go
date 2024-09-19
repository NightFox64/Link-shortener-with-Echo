package model

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrIDExist  = errors.New("this id already exists")
)

type AllURLModel struct {
	OriginalURL string
	ShortURL    string `gorm:"unique_short_index"`
}
