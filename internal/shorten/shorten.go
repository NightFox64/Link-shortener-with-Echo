package shorten

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
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

func PrependBaseURL(baseURL, identifier string) (string, error) {
	parsed, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	parsed.Path = identifier

	return parsed.String(), nil
}
