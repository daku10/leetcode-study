package main

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {

	testcases := []struct {
		arg  []int
		want int
	}{
		{
			[]int{2, 2, 3, 2},
			3,
		},
		{
			[]int{0, 1, 0, 1, 0, 1, 99},
			99,
		},
		{
			[]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2},
			-4,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {

			actual := singleNumber(tc.arg)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
