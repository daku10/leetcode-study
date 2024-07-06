package main

import (
	"fmt"
	"testing"
)

func TestMaxConsecutiveOnes(t *testing.T) {
	testcases := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 1, 0, 1, 1, 1},
			3,
		},
		{
			[]int{1, 0, 1, 1, 0, 1},
			2,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{0},
			0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findMaxConsecutiveOnes(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
