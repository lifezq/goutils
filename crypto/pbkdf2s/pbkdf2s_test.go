// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package pbkdf2

import (
	"fmt"
	"testing"
)

func TestPbkdf2Hash(t *testing.T) {

	for i := 0; i < 10; i++ {

		password := []byte(fmt.Sprintf("%d.%s", i, "abcdef@123"))

		h, err := GenerateFromPassword(password, 64)
		if err != nil {
			t.Errorf("Hash Error : %s \n", err.Error())
		}

		if !CompareHashAndPassword(h, password) {
			t.Errorf("CompareHashAndPassword Not Match. hashed:%v password:%v\n", h, password)
		}
	}
}
