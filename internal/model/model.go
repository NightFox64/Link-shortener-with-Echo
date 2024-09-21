package model

type AllURLModel struct {
	OriginalURL string
	ShortURL    string `gorm:"unique_short_index"`
}
