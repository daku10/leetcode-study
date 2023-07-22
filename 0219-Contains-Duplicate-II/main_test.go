package main

import (
	"fmt"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	testcases := []struct {
		nums []int
		k    int
		want bool
	}{
		{
			[]int{1, 2, 3, 1},
			3,
			true,
		},
		{
			[]int{1, 0, 1, 1},
			1,
			true,
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			2,
			false,
		},
		{
			[]int{99, 99},
			2,
			true,
		},
		{
			[]int{99, 98, 99},
			1,
			false,
		},
		{
			[]int{1, 2, 2, 3},
			3,
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := containsNearbyDuplicate(tc.nums, tc.k)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
