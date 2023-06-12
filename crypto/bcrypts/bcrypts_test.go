// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package bcrypts

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGeneratePassword(t *testing.T) {
	password := []byte("abcdefg@123")
	hashp, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		t.Errorf("Generate Password Error : %s\n", err.Error())
	}

	if err = bcrypt.CompareHashAndPassword(hashp, password); err != nil {
		t.Errorf("CompareHashAndPassword Error :%s \n", err.Error())
	}

	cost, err := bcrypt.Cost(hashp)
	if err != nil {
		t.Errorf("Cost Error:%s\n", err.Error())
	}

	t.Logf("Hashed Password Cost:%d  Len:%d : %s\n", cost, len(hashp), hashp)
}
