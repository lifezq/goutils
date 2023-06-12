// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package url

import (
	"fmt"
	"net/url"
	"strings"
)

var domainNSLen = 2

func ParseHostWithScheme(s string) (string, error) {
	urlObj, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s://%s", urlObj.Scheme, urlObj.Host), nil
}

func GetRootDomain(domain string) string {
	ns := strings.Split(domain, ":")
	ns = strings.Split(ns[0], ".")
	if len(ns) < domainNSLen {
		return ns[0]
	}

	return fmt.Sprintf(".%s.%s", ns[len(ns)-2], ns[len(ns)-1])
}
