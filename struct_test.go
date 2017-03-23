// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

import (
	"reflect"
	"testing"
	"unsafe"
)

///  /*
///   * basic types
///   */
///  typedef signed char             int8;
///  typedef unsigned char           uint8;
///  typedef signed short            int16;
///  typedef unsigned short          uint16;
///  typedef signed int              int32;
///  typedef unsigned int            uint32;
///  typedef signed long long int    int64;
///  typedef unsigned long long int  uint64;
///  typedef float                   float32;
///  typedef double                  float64;
///
///  #ifdef _64BIT
///  typedef uint64          uintptr;
///  typedef int64           intptr;
///  typedef int64           intgo; // Go's int
///  typedef uint64          uintgo; // Go's uint
///  #else
///  typedef uint32          uintptr;
///  typedef int32           intptr;
///  typedef int32           intgo; // Go's int
///  typedef uint32          uintgo; // Go's uint
///  #endif
///
///  /*
///   * defined types
///   */
///  typedef uint8           bool;
///  typedef uint8           byte;//

// byte
// output:uint8
func TestByte(t *testing.T) {
	var b byte = '0'
	t.Logf("ByteTest output:%v", reflect.TypeOf(b).Kind())
}

// rune
// output:int32
func TestRune(t *testing.T) {
	var r rune = 0
	t.Logf("RuneTest output:%v", reflect.TypeOf(r).Kind())
}

// string
//
// struct String
// {
//    byte* str;
//    intgo len;
// };
//
// output:&{5311546 12} true
func TestString(t *testing.T) {
	var str string = "Hello world!"

	p := (*struct {
		str  uintptr
		_len int
	})(unsafe.Pointer(&str))

	src_str := (*string)(unsafe.Pointer(p))

	t.Logf("TestString output:%v %v", p, *src_str == str)
}

// slice
//
// struct  Slice
// {					  // must not move anything
//    byte*   array;      // actual data
//    uintgo  len;        // number of elements
//    uintgo  cap;        // allocated number of elements
// };
//
// output:&{842350536768 5 10}
func TestSlice(t *testing.T) {
	var slice []int32 = make([]int32, 5, 10)

	p := (*struct {
		array uintptr
		_len  int
		_cap  int
	})(unsafe.Pointer(&slice))

	t.Logf("TestSlice output:%v", p)
}
