package url

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var charRunes = []rune(
	"1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateShort(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = charRunes[rand.Intn(len(charRunes))]
	}
	return string(b)
}
