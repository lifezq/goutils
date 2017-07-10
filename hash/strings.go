package hash

import (
	"math/rand"
	"time"
)

const (
	encode32Std = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	encode64Std = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func Rand32String(length uint8) string {

	rand.Seed(time.Now().UnixNano())

	b := []byte{}
	for len(b) < int(length) {
		b = append(b, encode32Std[rand.Intn(32)])
	}

	return string(b)
}

func Rand64String(length uint8) string {

	rand.Seed(time.Now().UnixNano())

	b := []byte{}
	for len(b) < int(length) {
		b = append(b, encode64Std[rand.Intn(64)])
	}

	return string(b)
}
