package utils

import (
	"errors"
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

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = int64(len(alphabet))
)

var (
	errZeroLength         = errors.New("Error zero length")
	errIncorrectCharacter = errors.New("Error: unexpected character in base62 literal")
)

//// EncodeBase62 ...
// func EncodeBase62(n int64) (string, error) {
// 	if n == 0 {
// 		return "", errZeroLength
// 	}
// 	buffer := make([]byte, 512)
// 	for n > 0 {
// 		r := math.Mod(float64(n), float64(base))
// 		n = n / base
// 		buffer = append([]byte{alphabet[int(r)]}, buffer...)
// 	}
// 	return string(buffer), nil
// }

// // DecodeBase62 ...
// func DecodeBase62(str string) (int64, error) {
// 	var r int64
// 	for _, c := range []byte(str) {
// 		i := strings.IndexByte(alphabet, c)
// 		if i < 0 {
// 			return 0, errIncorrectCharacter
// 		}
// 		r = base*r + int64(i)
// 	}
// 	return r, nil
// }
