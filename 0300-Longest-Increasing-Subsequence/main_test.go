package main

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	testcases := []struct {
		nums []int
		want int
	}{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{0, 1, 0, 3, 2, 3},
			4,
		},
		{
			[]int{7, 7, 7, 7, 7, 7, 7},
			1,
		},
		{
			[]int{4, 10, 4, 3, 8, 9},
			3,
		},
		{
			[]int{0, 1, 0, 3, 2, 3},
			4,
		},
		{
			[]int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			6,
		},
		{
			[]int{0},
			1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := lengthOfLIS(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
