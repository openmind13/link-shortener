package utils_test

import (
	"testing"

	"github.com/openmind13/link-shortener/app/utils"
	"github.com/stretchr/testify/assert"
)

func Test_RandStringRunes(t *testing.T) {
	count := 7
	assert.Equal(t, count, len(utils.GenerateRandomShortURL(count)))
}

func Test_ValidateURL(t *testing.T) {
	testCases := []struct {
		name    string
		url     string
		isValid bool
	}{
		{
			name:    "no root domain",
			url:     "https://vk.com/im",
			isValid: true,
		},
		{
			name:    "all is correct",
			url:     "https://youtube.com",
			isValid: true,
		},
		{
			name:    "no protocol",
			url:     "github.com",
			isValid: true,
		},
		{
			name:    "invalid characters",
			url:     "$$$example.com/fi///le[/].html",
			isValid: false,
		},
		{
			name:    "invalid format",
			url:     "abba",
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, utils.ValidateURL(tc.url))
			} else {
				assert.Error(t, utils.ValidateURL(tc.url))
			}
		})
	}
}
