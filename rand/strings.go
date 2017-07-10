package rand

import (
	"math/rand"
	"time"
)

const (
	std32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	std64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Rand32String(length uint8) string {

	b := []byte{}
	for len(b) < int(length) {
		b = append(b, std32[rand.Intn(32)])
	}

	return string(b)
}

func Rand64String(length uint8) string {

	b := []byte{}
	for len(b) < int(length) {
		b = append(b, std64[rand.Intn(64)])
	}

	return string(b)
}
