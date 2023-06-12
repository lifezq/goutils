// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package path

import (
	"testing"

	"github.com/lifezq/goutils/testings"
)

func TestIsFilename(t *testing.T) {
	data := testings.NewHaveWant("/abc/test", true)
	if !data.Equal(IsFilename(data.Have.(string))) {
		t.Error(data.Message)
	}

	data = testings.NewHaveWant("/abc/test/", false)
	if !data.Equal(IsFilename(data.Have.(string))) {
		t.Error(data.Message)
	}
}

func TestGetFilePath(t *testing.T) {
	data := testings.NewHaveWant("/abc/test/name", "abc/test")
	if !data.Equal(GetFilePath(data.Have.(string))) {
		t.Error(data.Message)
	}
}

func TestValidFilePath(t *testing.T) {
	data := testings.NewHaveWantMessage("/abc/test/name", true, "not valid file path")
	if !data.Equal(ValidFilePath(data.Have.(string))) {
		t.Error(data.Message)
	}
}
