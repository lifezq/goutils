// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package math

import "testing"

func TestLimits(t *testing.T) {
	if int8max := int8((0x01 << 7) ^ (^uint8(0))); Int8Max != int8max {
		t.Errorf("Int8Max not match %v-%v\n", Int8Max, int8max)
	}

	if uint8max := ^uint8(0); Uint8Max != uint8max {
		t.Errorf("Uint8Max not match %v-%v\n", Uint8Max, uint8max)
	}

	if int16max := int16((0x01 << 15) ^ (^uint16(0))); Int16Max != int16max {
		t.Errorf("Int16Max not match %v-%v\n", Int16Max, int16max)
	}

	if uint16max := ^uint16(0); Uint16Max != uint16max {
		t.Errorf("Uint16Max not match %v-%v\n", Uint16Max, uint16max)
	}

	if int32max := int32((0x01 << 31) ^ (^uint32(0))); Int32Max != int32max {
		t.Errorf("Int32Max not match %v-%v\n", Int32Max, int32max)
	}

	if uint32max := ^uint32(0); Uint32Max != uint32max {
		t.Errorf("Uint32Max not match %v-%v\n", Uint32Max, uint32max)
	}

	if int64max := int64((0x01 << 63) ^ (^uint64(0))); Int64Max != int64max {
		t.Errorf("Int64Max not match %v-%v\n", Int64Max, int64max)
	}

	if uint64max := ^uint64(0); Uint64Max != uint64max {
		t.Errorf("Uint64Max not match %v-%v\n", Uint64Max, uint64max)
	}
}
