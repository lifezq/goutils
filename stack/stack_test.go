// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package stack

import (
	"reflect"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    Stack[T]
		want bool
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
		{
			name: "空集测试",
			s:    Stack[int]{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    Stack[T]
		want T
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
		{
			name: "Pop测试",
			s:    Stack[int]{data: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args[T any] struct {
		v T
	}
	type testCase[T any] struct {
		name string
		s    Stack[T]
		args args[T]
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
		{
			name: "Pop测试",
			s:    Stack[int]{data: []int{1, 2, 3, 4, 5}},
			args: args[int]{6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.v)
		})
	}
}
