// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"encoding/base64"
	"fmt"
)

func Base64EncodeToString(src []byte) string {

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(buf, src)

	return string(buf)
}

func Base64DecodeString(s string) ([]byte, error) {

	dlen := base64.StdEncoding.DecodedLen(len([]byte(s)))
	buf := make([]byte, dlen)
	n, err := base64.StdEncoding.Decode(buf, []byte(s))
	if err != nil {
		return nil, err
	}

	if n != dlen {
		return buf, fmt.Errorf("Decode length not match")
	}

	return buf, nil
}
