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
