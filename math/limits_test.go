package math

import "testing"

func TestLimits(t *testing.T) {

	if int8max := int8((0x01 << 7) ^ (^uint8(0))); INT8_MAX != int8max {
		t.Errorf("INT8_MAX not match %v-%v\n", INT8_MAX, int8max)
	}

	if uint8max := ^uint8(0); UINT8_MAX != uint8max {
		t.Errorf("UINT8_MAX not match %v-%v\n", UINT8_MAX, uint8max)
	}

	if int16max := int16((0x01 << 15) ^ (^uint16(0))); INT16_MAX != int16max {
		t.Errorf("INT16_MAX not match %v-%v\n", INT16_MAX, int16max)
	}

	if uint16max := ^uint16(0); UINT16_MAX != uint16max {
		t.Errorf("UINT16_MAX not match %v-%v\n", UINT16_MAX, uint16max)
	}

	if int32max := int32((0x01 << 31) ^ (^uint32(0))); INT32_MAX != int32max {
		t.Errorf("INT32_MAX not match %v-%v\n", INT32_MAX, int32max)
	}

	if uint32max := ^uint32(0); UINT32_MAX != uint32max {
		t.Errorf("UINT32_MAX not match %v-%v\n", UINT32_MAX, uint32max)
	}

	if int64max := int64((0x01 << 63) ^ (^uint64(0))); INT64_MAX != int64max {
		t.Errorf("INT64_MAX not match %v-%v\n", INT64_MAX, int64max)
	}

	if uint64max := ^uint64(0); UINT64_MAX != uint64max {
		t.Errorf("UINT64_MAX not match %v-%v\n", UINT64_MAX, uint64max)
	}
}
