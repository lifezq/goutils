// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package cookie

import "testing"

func TestGenerateCookieSignValue(t *testing.T) {
	value, err := GenerateCookieSignValue("cvalue")
	if err != nil {
		t.Error(err)
	}

	item := NewCookieWithDomain("ctest", value, "localhost")
	_, ok := SignVerify(item)
	if !ok {
		t.Error("Unexpected")
	}
}
