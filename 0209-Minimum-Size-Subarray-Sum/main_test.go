package main

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	testcases := []struct {
		target int
		nums   []int
		want   int
	}{
		{
			7,
			[]int{2, 3, 1, 2, 4, 3},
			2,
		},
		{
			4,
			[]int{1, 4, 4},
			1,
		},
		{
			11,
			[]int{1, 1, 1, 1, 1, 1, 1, 1},
			0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minSubArrayLen(tc.target, tc.nums)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
