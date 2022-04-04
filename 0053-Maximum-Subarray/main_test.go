package main

import (
	"fmt"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {

	testcase := []struct {
		arg  []int
		want int
	}{
		{
			arg:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			arg:  []int{1},
			want: 1,
		},
		{
			arg:  []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			arg:  []int{-1},
			want: -1,
		},
		{
			arg:  []int{-1, -2, -3, -4},
			want: -1,
		},
	}

	for _, tc := range testcase {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := maxSubArray(tc.arg)
			if tc.want != actual {
				t.Errorf("expect: %d, actual: %d\n", tc.want, actual)
			}
		})
	}
}
