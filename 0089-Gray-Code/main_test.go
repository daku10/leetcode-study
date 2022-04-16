package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	testcases := []struct {
		arg  int
		want []int
	}{
		{
			// 0 1
			1, []int{0, 1},
		},
		{
			// 00 01 11 10
			2, []int{0, 1, 3, 2},
		},
		{
			// 000 001 011 010 110 111 101 100
			3, []int{0, 1, 3, 2, 6, 7, 5, 4},
		},
		{
			// 0000 0001 0011 0010 0110 0111 0101 0100 1100 1101 1111 1110 1010 1011 1001 1000
			4, []int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := grayCode(tc.arg)
			if !reflect.DeepEqual(tc.want, actual) {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
