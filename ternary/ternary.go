// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package ternary

func IfSet[T any](condition bool, v1, v2 T) T {
	if condition {
		return v1
	}
	return v2
}
