// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package regexps

import "regexp"

var (
	DomainPreg   = regexp.MustCompile(`^[0-9a-zA-Z]+[0-9a-zA-Z.-]*\.[a-zA-Z]{2,4}$`)
	EmailPreg    = regexp.MustCompile(`^(\w)+(\.\w+)*@(\w)+((\.\w{2,3}){1,3})$`)
	PhonePreg    = regexp.MustCompile(`^1[358]\d{9}$`)
	UsernamePreg = regexp.MustCompile(`^[a-zA-Z][0-9a-zA-Z_]{5,29}$`)
)

// CheckDomain 域名检查
func CheckDomain(domain string) bool {
	return DomainPreg.MatchString(domain)
}

// CheckEmail 邮箱检查
func CheckEmail(email string) bool {
	return EmailPreg.MatchString(email)
}

// CheckPhone 手机号检查
func CheckPhone(phone string) bool {
	return PhonePreg.MatchString(phone)
}

// CheckUsername 用户名检查
func CheckUsername(username string) bool {
	return UsernamePreg.MatchString(username)
}
