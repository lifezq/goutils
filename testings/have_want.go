// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package testings

type HaveWant struct {
	Have    any
	Want    any
	Message string
}

func NewHaveWant(have, want any) *HaveWant {
	return &HaveWant{
		Have:    have,
		Want:    want,
		Message: "unexpected",
	}
}

func NewHaveWantMessage(have, want any, message string) *HaveWant {
	return &HaveWant{
		Have:    have,
		Want:    want,
		Message: message,
	}
}

func (h *HaveWant) Equal(result any) bool {
	return result == h.Want
}
