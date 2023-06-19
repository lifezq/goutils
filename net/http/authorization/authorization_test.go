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
	testings.NewEntities(&testings.Entity{
		Name: "test",
		Args: GetRequestAuthorization(http.Header{
			"Authorization": []string{"Bearer X-Auth-Token"},
		}),
		Want: "X-Auth-Token",
	}).Run(t)
}
