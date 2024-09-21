package shorten

import (
	"encoding/base64"
	"fmt"
	"math/rand"
)

func GenerateNewURL() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(rand.Int())))
}

func Shorten(urlEncoded string) string {
	if len(urlEncoded) < 8 {
		return urlEncoded
	}
	return urlEncoded[:8]
}
