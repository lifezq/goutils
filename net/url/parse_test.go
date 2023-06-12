// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package url

import "testing"

func TestParseHostWithScheme(t *testing.T) {
	s := "http://localhost:8080/api/v1/users"
	host, err := ParseHostWithScheme(s)
	if err != nil {
		t.Error(err)
	}

	want := "http://localhost:8080"
	if host != want {
		t.Error("unexpected")
	}
}

func TestGetRootDomain(t *testing.T) {
	domain := "ryan.exe.life.com"
	want := ".life.com"
	if GetRootDomain(domain) != want {
		t.Error("error not matching")
	}

	domain = "localhost"
	want = "localhost"
	if GetRootDomain(domain) != want {
		t.Error("error not matching")
	}
}
