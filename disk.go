// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetDiskUUIDNumber(name string) string {

	n := fmt.Sprintf("--name=%s", name)
	udevadm, err := exec.LookPath("/sbin/udevadm")
	if err != nil {
		return ""
	}

	out, err := exec.Command(udevadm, "info", "--query=property", n).Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {

		values := strings.Split(line, "=")
		if len(values) < 2 || values[0] != "ID_FS_UUID" {
			// only get ID_FS_UUID
			continue
		}

		return values[1]
	}

	return ""
}

func GetRelativePath(p1, p2 string) string {

	p1arr := strings.Split(p1, "/")
	p2arr := strings.Split(p2, "/")

	minLen := len(p1arr)
	if len(p2arr) < minLen {
		minLen = len(p2arr)
	}

	diff := 0
	for diff <= minLen {
		if p1arr[diff] != p2arr[diff] {
			break
		}

		diff++
	}

	return strings.Repeat("../", len(p2arr)-diff-1) + strings.Join(p1arr[diff:], "/")
}
