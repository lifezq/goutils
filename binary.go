// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

// uint8  = 1<<8-1
// int8   = 1<<7-1
// uint16 = 1<<16-1
// int16  = 1<<15-1
// uint32 = 1<<32-1
// int32  = 1<<31-1
// uint64 = 1<<64-1
// int64  = 1<<63-1

func ZoreInt(n int) int {
	return n ^ n //nolint:gocritic,staticcheck	// ignore
}

func IsOddNumber(n int) bool {
	return n&0x01 != 0x01
}

// a|b&b == a
// a&b|b == a.
func IsEqual(a, b int) bool {
	return ((a | b) & b) == a
}

func BoolToInt8(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func Int8ToBool(i int8) bool {
	return i == 1
}

func EncryptNumber(n int) int {
	return n ^ 0x01<<24
}

func EncryptBytes(bs []byte) {
	for i, b := range bs {
		bs[i] = b ^ 0x01<<4
	}
}

func BigEndianUint16ToBytes(bs []byte, n uint16) {
	if cap(bs) < 2 {
		return
	}

	bs[0] = byte(n >> 8)
	bs[1] = byte(n)
}

func BigEndianBytesToUint16(bs []byte) uint16 {
	if cap(bs) < 2 {
		return 0
	}

	return uint16(bs[1]) | uint16(bs[0])<<8
}

func BigEndianUint32ToBytes(bs []byte, n uint32) {
	if cap(bs) < 4 {
		return
	}

	bs[0] = byte(n >> 24)
	bs[1] = byte(n >> 16)
	bs[2] = byte(n >> 8)
	bs[3] = byte(n)
}

func BigEndianBytesToUint32(bs []byte) uint32 {
	if cap(bs) < 4 {
		return 0
	}

	return uint32(bs[3]) | uint32(bs[2])<<8 | uint32(bs[1])<<16 | uint32(bs[0])<<24
}

func BigEndianUint64ToBytes(bs []byte, n uint64) {
	if cap(bs) < 8 {
		return
	}

	bs[0] = byte(n >> 56)
	bs[1] = byte(n >> 48)
	bs[2] = byte(n >> 40)
	bs[3] = byte(n >> 32)
	bs[4] = byte(n >> 24)
	bs[5] = byte(n >> 16)
	bs[6] = byte(n >> 8)
	bs[7] = byte(n)
}

func BigEndianBytesToUint64(bs []byte) uint64 {
	if cap(bs) < 8 {
		return 0
	}

	return uint64(bs[7]) | uint64(bs[6])<<8 | uint64(bs[5])<<16 | uint64(bs[4])<<24 |
		uint64(bs[3])<<32 | uint64(bs[2])<<40 | uint64(bs[1])<<48 | uint64(bs[0])<<56
}
