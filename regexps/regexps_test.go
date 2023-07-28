// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package regexps

import (
	"testing"

	"github.com/lifezq/goutils/testings"
)

func TestCheckDomain(t *testing.T) {
	entities := testings.Entities{
		{
			Name: "#00",
			Args: CheckDomain("www.example.com"),
			Want: true,
		},
		{
			Name: "#01",
			Args: CheckDomain("sub.face.example.com"),
			Want: true,
		},
	}

	entities.Run(t)
}

func TestCheckEmail(t *testing.T) {
	entities := testings.Entities{
		{
			Name: "#00",
			Args: CheckEmail("lifezqy@126.com"),
			Want: true,
		},
		{
			Name: "#01",
			Args: CheckEmail("yangqianleizq@gmail.com"),
			Want: true,
		},
		{
			Name: "#01",
			Args: CheckEmail("yangqianleizq#gmail.com"),
			Want: false,
		},
	}

	entities.Run(t)
}

func TestCheckPhone(t *testing.T) {
	entities := testings.Entities{
		{
			Name: "#00",
			Args: CheckPhone("15001056879"),
			Want: true,
		},
		{
			Name: "#01",
			Args: CheckPhone("35001056879"),
			Want: false,
		},
	}

	entities.Run(t)
}

func TestCheckUsername(t *testing.T) {
	entities := testings.Entities{
		{
			Name: "#00",
			Args: CheckUsername("Zush23_d"),
			Want: true,
		},
		{
			Name: "#01",
			Args: CheckUsername("2Uns.dks@"),
			Want: false,
		},
	}

	entities.Run(t)
}
