package model

import (
	"errors"

	. "github.com/samber/mo"
)

var (
	ErrNotFound = errors.New("not found")
	ErrIDExist  = errors.New("this id already exists")
)

type AllURLModel struct {
	Identifier  string `json:"identifier"`
	OriginalURL string
	ShortURL    string `gorm:"unique_short_index"`
}

type InputURLModel struct {
	RawURL     string
	Identifier Option[string]
}
