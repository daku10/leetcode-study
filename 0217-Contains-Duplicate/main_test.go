package main

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	testcases := []struct {
		nums []int
		want bool
	}{
		{
			[]int{1, 2, 3, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4},
			false,
		},
		{
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := containsDuplicate(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
