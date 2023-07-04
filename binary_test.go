// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

import (
	"strconv"
	"testing"
)

func TestBinaryZoreInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if ZoreInt(i) != 0 {
			t.Errorf("ZoreInt error:%d\n", i)
		}
	}
}

func TestBinaryIsOddNumber(t *testing.T) {
	for i := 0; i < 1000; i += 2 {
		if !IsOddNumber(i) {
			t.Errorf("IsOddNumber error:%d\n", i)
		}
	}
}

func TestBinaryIsEqual(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if IsEqual(i, j) != (i == j) {
				t.Errorf("IsEqual error:%d\n", i)
			}
		}
	}
}

func TestBoolToInt8(t *testing.T) {
	b := false
	if BoolToInt8(b) != 0 {
		t.Error("TestBoolToInt8 error")
	}

	b = true
	if BoolToInt8(b) != 1 {
		t.Error("TestBoolToInt8 error")
	}
}

func TestInt8ToBool(t *testing.T) {
	i := int8(0)
	if Int8ToBool(i) {
		t.Error("TestInt8ToBool error")
	}

	i = int8(1)
	if !Int8ToBool(i) {
		t.Error("TestInt8ToBool error")
	}
}

func TestBinaryEncryptNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if i != EncryptNumber(EncryptNumber(i)) {
			t.Errorf("EncryptNumber error:%d\n", i)
		}
	}
}

func TestBinaryEncryptBytes(t *testing.T) {
	s := "test encrypt bytes..."

	for i := 0; i < 1000; i++ {
		bs := []byte(s + strconv.Itoa(i))

		// encode
		EncryptBytes(bs)

		// decode
		EncryptBytes(bs)

		if string(bs) != s+strconv.Itoa(i) {
			t.Errorf("EncryptBytes error:%d\n", i)
		}
	}
}

func TestBinaryBigEndianBytesAndUint16(t *testing.T) {
	bs := make([]byte, 2)

	for i := uint16(0); i < 1000; i++ {
		BigEndianUint16ToBytes(bs, i)

		if BigEndianBytesToUint16(bs) != i {
			t.Errorf("BytesAndUint16 error:%d\n", i)
		}
	}
}

func TestBinaryBigEndianBytesAndUint32(t *testing.T) {
	bs := make([]byte, 4)

	for i := uint32(0); i < 1000; i++ {
		BigEndianUint32ToBytes(bs, i)

		if BigEndianBytesToUint32(bs) != i {
			t.Errorf("BytesAndUint32 error:%d\n", i)
		}
	}
}

func TestBinaryBigEndianBytesAndUint64(t *testing.T) {
	bs := make([]byte, 8)

	for i := uint64(0); i < 1000; i++ {
		BigEndianUint64ToBytes(bs, i)

		if BigEndianBytesToUint64(bs) != i {
			t.Errorf("BytesAndUint64 error:%d\n", i)
		}
	}
}
