// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import "testing"

func TestBase64EncodeAndDecode(t *testing.T) {

	str := "Encode test string..."
	estr := Base64EncodeToString([]byte(str))

	destr, err := Base64DecodeString(estr)
	if err != nil {
		t.Errorf("Base64DecodeString : %s\n", err.Error())
	}

	t.Logf("Decoded : %s\n", string(destr))
}
