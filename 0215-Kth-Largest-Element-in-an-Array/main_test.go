package main

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	testcases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			[]int{3, 2, 1, 5, 6, 4},
			2,
			5,
		},
		{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			4,
			4,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findKthLargest(tc.nums, tc.k)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
