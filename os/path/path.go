// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package path

import (
	"regexp"
	"strings"
)

var invalidPathPattern = regexp.MustCompile(`[>|<?@#$%^&*+]|(//)|(\\\\)`)

// IsFilename Determine if it is a file name.
func IsFilename(p string) bool {
	return !strings.HasSuffix(p, "/")
}

func trimPrefixSlash(p string) string {
	return strings.TrimPrefix(strings.Trim(p, " "), "/")
}

// GetFilePath Get file directory path.
func GetFilePath(p string) string {
	p = trimPrefixSlash(p)
	idx := strings.LastIndex(p, "/")
	if idx <= 0 {
		return ""
	}

	return p[:idx]
}

// ValidFilePath Verify if it is a valid file name path.
func ValidFilePath(p string) bool {
	return !invalidPathPattern.MatchString(p)
}
