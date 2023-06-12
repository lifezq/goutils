// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package authorization

import (
	"net/http"
	"testing"

	"github.com/lifezq/goutils/testings"
)

func TestGetRequestAuthorization(t *testing.T) {
	data := testings.NewHaveWant(http.Header{
		"Authorization": []string{"Bearer X-Auth-Token"},
	}, "X-Auth-Token")

	if !data.Equal(GetRequestAuthorization(data.Have.(http.Header))) {
		t.Error("got unexpected authorization")
	}
}
