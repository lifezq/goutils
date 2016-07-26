// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package regexps

import "regexp"

var (
	DomainPreg   = regexp.MustCompile(`^[0-9a-zA-Z]+[0-9a-zA-Z\.-]*\.[a-zA-Z]{2,4}$`)
	EmailPreg    = regexp.MustCompile(`^(\w)+(\.\w+)*@(\w)+((\.\w{2,3}){1,3})$`)
	PhonePreg    = regexp.MustCompile(`^[1][358]\d{9}$`)
	UsernamePreg = regexp.MustCompile(`^[a-zA-Z][0-9a-zA-Z_]{5,29}$`)
)
