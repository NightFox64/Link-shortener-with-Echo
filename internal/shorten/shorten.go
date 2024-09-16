package shorten

import (
	"encoding/base64"
)

func GenerateNewURL(url string) string {
	return base64.StdEncoding.EncodeToString([]byte(url))
}

func Shorten(urlEncoded string) string {
	if len(urlEncoded) < 8 {
		return urlEncoded
	}
	return urlEncoded[:8]
}
