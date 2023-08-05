// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package ternary

import (
	"testing"

	"github.com/lifezq/goutils/testings"
)

func TestIfSet(t *testing.T) {
	testings.NewEntities(&testings.Entity{
		Args: IfSet(true, "是的", "不是"),
		Want: "是的",
		Name: "#00",
	}, &testings.Entity{
		Args: IfSet(false, 1, 2),
		Want: 2,
		Name: "#01",
	}).Run(t)
}
