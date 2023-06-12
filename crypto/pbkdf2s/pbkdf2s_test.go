// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package pbkdf2s

import (
	"fmt"
	"testing"

	"github.com/lifezq/goutils/encode"
)

func TestPbkdf2GeneratePassword(t *testing.T) {
	for i := 0; i < 10; i++ {
		password := []byte(fmt.Sprintf("%d.%s", i, "abcdef@123"))

		h, err := GenerateFromPassword(password, 32)
		if err != nil {
			t.Errorf("GenerateFromPassword Error : %s \n", err.Error())
		}

		if !CompareHashAndPassword(h, password) {
			t.Errorf("CompareHashAndPassword Not Match. hashed:%v password:%v\n", h, password)
		}

		t.Logf("Pbkdf2 hashed len:%d hash:%v hex:%s\n", len(h), h, encode.HexEncodeToString(h))
	}
}

func TestScryptGeneratePassword(t *testing.T) {
	for i := 0; i < 10; i++ {
		password := []byte(fmt.Sprintf("%d.%s", i, "abcdef@123"))

		h, err := ScryptGenerateFromPassword(password, 32)
		if err != nil {
			t.Errorf("ScryptGenerateFromPassword Error : %s \n", err.Error())
		}

		if !ScryptCompareHashAndPassword(h, password) {
			t.Errorf("ScryptCompareHashAndPassword Not Match. hashed:%v password:%v\n", h, password)
		}

		t.Logf("Scrypt hashed len:%d hash:%v hex:%s\n", len(h), h, encode.HexEncodeToString(h))
	}
}
