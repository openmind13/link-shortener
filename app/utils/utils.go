package utils

import (
	"math/rand"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GenerateRandomShortURL - return random string with "n" length
func GenerateRandomShortURL(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

// ValidateURL - check correct url or not
func ValidateURL(url string) error {
	return validation.Validate(
		url,
		// rules
		validation.Required,
		is.URL,
	)
}
