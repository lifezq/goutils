// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package rand

import (
	"math/rand"
	"strconv"
)

//nolint:gosec // ignore
func DigitalVerificationCode(length int) string {
	n := ""
	for i := 0; i < length; i++ {
		n += strconv.Itoa(rand.Intn(10))
	}
	return n
}
