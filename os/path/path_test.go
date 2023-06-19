// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package path

import (
	"testing"

	"github.com/lifezq/goutils/testings"
)

func TestIsFilename(t *testing.T) {
	testings.NewEntities(
		&testings.Entity{
			Name: "test1",
			Args: IsFilename("/abc/test"),
			Want: true,
		}, &testings.Entity{
			Name: "test2",
			Args: IsFilename("/abc/test/"),
			Want: false,
		}).Run(t)
}

func TestGetFilePath(t *testing.T) {
	testings.NewEntities(&testings.Entity{
		Name: "test3",
		Args: GetFilePath("/abc/test/name"),
		Want: "abc/test",
	}).Run(t)
}

func TestValidFilePath(t *testing.T) {
	testings.NewEntities(&testings.Entity{
		Name: "test4",
		Args: ValidFilePath("/abc/test/name"),
		Want: true,
	}).Run(t)
}
