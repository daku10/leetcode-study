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
			[]int{2, 3, 2},
			3,
		},
		{
			[]int{1, 2, 3, 1},
			4,
		},
		{
			[]int{
				1, 2, 3,
			},
			3,
		},
		{
			[]int{1, 3, 1, 3, 100},
			103,
		},
		{
			[]int{4, 1, 2, 7, 5, 3, 1},
			14,
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
