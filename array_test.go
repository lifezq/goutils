// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

import "testing"

func TestArrayRemoveDup(t *testing.T) {

	for i := 0; i < 100; i++ {

		arr := []interface{}{1, 2, 5, 6, 11, 25, 33, 66, 15, 25, 15, 66}

		striped := ArrayRemoveDup(arr)
		if len(striped) != 9 {
			t.Errorf("ArrayRemoveDup error\n")
		}

		//t.Logf("ArrayRemoveDup source:%v striped:%v\n", arr, len(striped))
	}
}
