// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package cookie

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SignValue struct {
	IssuerAt string
	Data     string
	Sign     string
}

const cookieSalt = "salt:secret"

var defaultMaxAge = 86400

func NewCookieWithDomain(name, value, domain string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   domain,
		MaxAge:   defaultMaxAge,
		Expires:  time.Now().Add(time.Duration(defaultMaxAge) * time.Second),
		HttpOnly: false,
	}
}

func NewCookieWithDomainAndMaxAge(name, value, domain string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   domain,
		MaxAge:   maxAge,
		Expires:  time.Now().Add(time.Duration(maxAge) * time.Second),
		HttpOnly: false,
	}
}

func GenerateCookieSignValue(data string) (string, error) {
	now := time.Now().String()
	value := &SignValue{
		IssuerAt: now,
		Data:     data,
		Sign:     sign(now + data),
	}

	jsp, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(jsp), nil
}

func SignVerify(data *http.Cookie) (string, bool) {
	err := data.Valid()
	if err != nil {
		return "", false
	}

	item, err := base64.RawStdEncoding.DecodeString(data.Value)
	if err != nil {
		return "", false
	}

	var value SignValue
	err = json.Unmarshal(item, &value)
	if err != nil {
		return "", false
	}

	if sign(value.IssuerAt+value.Data) != value.Sign {
		return "", false
	}

	return value.Data, true
}

func sign(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum([]byte(cookieSalt)))
}
