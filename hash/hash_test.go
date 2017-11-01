// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package hash

import "testing"

func TestStringHash(t *testing.T) {

	wanted := StringHash("hash.test", 16)
	for i := 0; i < 100; i++ {
		if got := StringHash("hash.test", 16); got != wanted {
			t.Fatalf("StringHash error wanted:%s got:%s\n", wanted, got)
		}
	}
}

func TestSaltHash(t *testing.T) {

	wanted := SaltHash("hash.test", "salt", 16)
	for i := 0; i < 100; i++ {
		if got := SaltHash("hash.test", "salt", 16); got != wanted {
			t.Fatalf("StringHash error wanted:%s got:%s\n", wanted, got)
		}
	}
}
