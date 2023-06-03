package main

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	testcases := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 2, 3, 1},
			4,
		},
		{
			[]int{2, 1, 1, 3},
			5,
		},
		{
			[]int{2, 7, 9, 3, 1},
			12,
		},
		{
			[]int{2, 7, 9, 3, 1, 15},
			26,
		},
		{
			[]int{2, 4, 8, 9, 9, 3},
			19,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 2, 3},
			4,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := rob(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
