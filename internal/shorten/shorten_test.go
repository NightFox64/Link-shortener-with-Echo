package shorten_test

import (
	"testing"

	"github.com/NightFox64/Link-shortener-with-Echo/internal/shorten"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	t.Run("returns a short URL", func(t *testing.T) {
		type testCase struct {
			url    string
			answer string
		}

		testCases := []testCase{
			{
				url:    "1234",
				answer: "MTIzNA==",
			},
			{
				url:    "1389qqjqio728190",
				answer: "MTM4OXFx",
			},
		}

		for _, tc := range testCases {
			actual := shorten.GenerateNewURL(tc.url)
			actual = shorten.Shorten(actual)
			assert.Equal(t, tc.answer, actual)
		}
	})
}
