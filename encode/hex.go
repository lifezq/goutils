// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"encoding/hex"
	"fmt"
)

func HexEncodeToString(b []byte) string {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	return string(dst)
}

func HexDecodeString(s string) ([]byte, error) {

	var (
		src  = []byte(s)
		dlen = hex.DecodedLen(len(src))
		dst  = make([]byte, dlen)
	)

	n, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}

	if n != dlen {
		return nil, fmt.Errorf("Wrong:copy length not match")
	}

	return dst, nil
}
