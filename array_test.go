// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

import "testing"

func TestArrayInsertRemove(t *testing.T) {

	arr := &Array{}
	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			arr.Insert(i)
		}
	}

	if len(*arr) != 100 {
		t.Errorf("ArrayInsert %s", "#01")
	}

	for i := 0; i < 100; i++ {
		arr.Remove(i)
	}

	if len(*arr) != 0 {
		t.Errorf("ArrayRemove %s", "#02")
	}
}

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
