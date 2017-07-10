package hash

import (
	"crypto/md5"
	"fmt"
)

func StringHash(s string, length uint16) string {

	if length < 1 {
		length = 1
	} else if length > 32 {
		length = 32
	}

	md5_val := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	return md5_val[:length]
}