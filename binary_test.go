// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

import (
	"strconv"
	"testing"
)

func TestZoreInt(t *testing.T) {

	for i := 0; i < 1000; i++ {
		if ZoreInt(i) != 0 {
			t.Errorf("ZoreInt error:%d\n", i)
		}
	}
}

func TestIsOddNumber(t *testing.T) {

	for i := 0; i < 1000; i += 2 {
		if !IsOddNumber(i) {
			t.Errorf("IsOddNumber error:%d\n", i)
		}
	}
}

func TestIsEqual(t *testing.T) {

	for i := 0; i < 1000; i++ {

		if !IsEqual(i, i) {
			t.Errorf("IsEqual error:%d\n", i)
		}
	}
}

func TestEncryptNumber(t *testing.T) {

	for i := 0; i < 1000; i++ {
		if i != EncryptNumber(EncryptNumber(i)) {
			t.Errorf("EncryptNumber error:%d\n", i)
		}
	}
}

func TestEncryptBytes(t *testing.T) {

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

func TestBigEndianBytesAndUint16(t *testing.T) {

	bs := make([]byte, 2)

	for i := uint16(0); i < 1000; i++ {

		BigEndianUint16ToBytes(bs, i)

		if BigEndianBytesToUint16(bs) != i {
			t.Errorf("BytesAndUint16 error:%d\n", i)
		}
	}
}

func TestBigEndianBytesAndUint32(t *testing.T) {

	bs := make([]byte, 4)

	for i := uint32(0); i < 1000; i++ {

		BigEndianUint32ToBytes(bs, i)

		if BigEndianBytesToUint32(bs) != i {
			t.Errorf("BytesAndUint32 error:%d\n", i)
		}
	}
}

func TestBigEndianBytesAndUint64(t *testing.T) {

	bs := make([]byte, 8)

	for i := uint64(0); i < 1000; i++ {

		BigEndianUint64ToBytes(bs, i)

		if BigEndianBytesToUint64(bs) != i {
			t.Errorf("BytesAndUint64 error:%d\n", i)
		}
	}
}
