package main

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	testcases := []struct {
		arg  []int
		want int
	}{
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
		{
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			9,
		},
		{
			[]int{1, 2, 0, 1},
			3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := longestConsecutive(tc.arg)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
