package goutils

import "testing"

func TestArrayRemoveDup(t *testing.T) {

	arr := []interface{}{1, 2, 5, 6, 11, 25, 33, 66, 15, 25, 15, 66}
	t.Logf("ArrayRemoveDup source:%v triped:%v\n", arr, ArrayRemoveDup(arr))
}
