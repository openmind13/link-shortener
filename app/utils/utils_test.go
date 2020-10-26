package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStringRunes(t *testing.T) {
	count := 7
	assert.Equal(t, count, len(GenerateRandomShortURL(count)))
}

// func TestBase62EncodeURL(t *testing.T) {
// 	count := 7
// 	hashStr := Base62EncodeURL(count)
// 	assert.Equal(t, count, len(hashStr))
// 	fmt.Println(hashStr)
// }

func TestValidateURL(t *testing.T) {
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
				assert.NoError(t, ValidateURL(tc.url))
			} else {
				assert.Error(t, ValidateURL(tc.url))
			}
		})
	}
}
