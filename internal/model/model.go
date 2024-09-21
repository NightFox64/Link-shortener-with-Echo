package model

type AllURLModel struct {
	Identifier  string `json:"identifier"`
	OriginalURL string
	ShortURL    string `gorm:"unique_short_index"`
}
