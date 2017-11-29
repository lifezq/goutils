// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
)

func StringHash(s string, length uint16) string {

	if length < 4 {
		length = 4
	} else if length > 32 {
		length = 32
	}

	md5_val := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	return md5_val[:length]
}

func SaltHash(s, salt string, length uint16) string {

	if length < 8 {
		length = 8
	} else if length > 64 {
		length = 64
	}

	h := hmac.New(sha256.New, []byte(salt))
	io.WriteString(h, s)

	md5_val := fmt.Sprintf("%x", h.Sum(nil))
	return md5_val[:length]
}
