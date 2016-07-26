// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package utils

func ArrayStringContain(a []string, s string) bool {

	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

func ArrayIntContain(a []int, s int) bool {

	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}
