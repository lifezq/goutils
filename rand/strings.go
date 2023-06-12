// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package rand

import (
	"math/rand"
	"time"
)

const (
	std32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	std64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // ignore
}

func Rand32String(length uint8) string {
	b := []byte{}
	for len(b) < int(length) {
		b = append(b, std32[rand.Intn(32)]) //nolint:gosec // ignore
	}

	return string(b)
}

func Rand64String(length uint8) string {
	b := []byte{}
	for len(b) < int(length) {
		b = append(b, std64[rand.Intn(64)]) //nolint:gosec // ignore
	}

	return string(b)
}
