// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

func ArrayStringContain(a []string, s string) bool {

	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

func arrayStringRemove(a []string, s string) []string {

	for i, v := range a {

		if v == s {
			a = append(a[:i], a[i+1:]...)
		}
	}

	return a
}

func ArrayIntContain(a []int, s int) bool {

	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

func arrayIntRemove(a []int, s int) []int {

	for i, v := range a {

		if v == s {
			a = append(a[:i], a[i+1:]...)
		}
	}

	return a
}
