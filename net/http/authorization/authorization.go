// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package authorization

import (
	"net/http"
	"strings"
)

const (
	BearerToken = "Bearer"

	PrefixBearer = BearerToken + " "
)

func GetRequestAuthorization(header http.Header) string {
	var auth = header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, PrefixBearer) {
		return ""
	}
	return strings.TrimPrefix(auth, PrefixBearer)
}
