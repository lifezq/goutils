// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

type Array []interface{}

func (a *Array) Contain(val interface{}) bool {

	for _, v := range *a {
		if v == val {
			return true
		}
	}

	return false
}

func (a *Array) Insert(val interface{}) {

	for _, v := range *a {
		if v == val {
			return
		}
	}

	*a = append(*a, val)
}

func (a *Array) Remove(val interface{}) {

	for i, v := range *a {
		if v == val {
			*a = append((*a)[:i], (*a)[i+1:]...)
			break
		}
	}
}

func ArrayStringContain(a []string, s string) bool {

	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

func ArrayStringRemove(a []string, s string) []string {

	for i, v := range a {

		if v == s {
			a = append(a[:i], a[i+1:]...)
		}
	}

	return a
}

func ArrayIntContain(a []int, n int) bool {

	for _, v := range a {
		if v == n {
			return true
		}
	}

	return false
}

func ArrayIntRemove(a []int, n int) []int {

	for i, v := range a {

		if v == n {
			a = append(a[:i], a[i+1:]...)
		}
	}

	return a
}

func ArrayRemoveDup(a []interface{}) []interface{} {

	if len(a) <= 1 {
		return a
	}

	var set []interface{}
	for _, n := range a {

		found := false
		for _, v := range set {

			if n == v {
				found = true
				break
			}
		}

		if found {
			continue
		}

		set = append(set, n)
	}

	return set
}
