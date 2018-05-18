// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package ciphers

import (
	"fmt"
	"testing"
)

func TestCBCEnDecrypt(t *testing.T) {

	var (
		encrypted string
		decrypted string
		err       error
	)

	for i := 0; i < 100; i++ {

		plaintext := fmt.Sprintf("%d.%s", i, "exampleplaintext")
		plaintext = plaintext[:16]

		encrypted, err = CBCEncrypt([]byte(plaintext))
		if err != nil {
			t.Errorf("CBCEncrypt error:%s\n", err.Error())
			continue
		}

		decrypted, err = CBCDecrypt(encrypted)
		if err != nil {
			t.Errorf("CBCDecrypt error:%s\n", err.Error())
			continue
		}

		if decrypted != plaintext {
			t.Errorf("CBC Encrypt/Decrypt wrong. plaintext:%s decrypted:%s\n",
				plaintext, decrypted)
		}
	}

	t.Logf("encrypt len:%d :%s decrypted:%s err:%v\n",
		len(encrypted), encrypted, decrypted, err)
}

func TestCFBEnDecrypter(t *testing.T) {

	var (
		encrypted string
		decrypted string
		err       error
	)

	for i := 0; i < 100; i++ {

		plaintext := fmt.Sprintf("%d.%s", i, "some plaintext any length")

		encrypted, err = CFBEncrypt([]byte(plaintext))
		if err != nil {
			t.Errorf("CFBEncrypt error:%s\n", err.Error())
			continue
		}

		decrypted, err = CFBDecrypt(encrypted)
		if err != nil {
			t.Errorf("CFBDecrypt error:%s\n", err.Error())
			continue
		}

		if decrypted != plaintext {
			t.Errorf("CFB Encrypt/Decrypt wrong. plaintext:%s decrypted:%s\n",
				plaintext, decrypted)
		}
	}

	t.Logf("encrypt len:%d :%s decrypted:%s err:%v\n",
		len(encrypted), encrypted, decrypted, err)
}
