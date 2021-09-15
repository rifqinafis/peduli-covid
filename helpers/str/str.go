package str

import (
	"math/rand"
	"time"
)

// RandAlphanumericString ...
func RandAlphanumericString(length int) string {
	charset := "0123456789"
	return StringWithCharset(length, charset)
}

// StringWithCharset ...
func StringWithCharset(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
