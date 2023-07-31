// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package rand

import "testing"

func TestDigitalVerificationCode(t *testing.T) {
	length := 6
	for i := 0; i < 1000; i++ {
		code := DigitalVerificationCode(length)
		if len(code) != length {
			t.Fatal("Invalid code length")
		}
	}
}
