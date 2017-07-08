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
