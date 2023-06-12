// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package rand

import (
	"crypto/rand"
	"fmt"

	mrand "math/rand"
)

// UUID Version 4.
func UUIDGen() string {
	var b [16]byte
	length := len(b)
	n, err := rand.Read(b[:])
	if n != length || err != nil {
		for length > 0 {
			length--
			b[length] = byte(mrand.Intn(256)) //nolint:gosec // ignore
		}
	}

	b[6] = (b[6] & 0x0f) | 4<<4 // Version 4
	b[8] = (b[8] & 0x3f) | 8<<4 // Variant 10

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}
