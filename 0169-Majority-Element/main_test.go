package main

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testcases := []struct {
		nums []int
		want int
	}{
		{
			[]int{3, 2, 3},
			3,
		},
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
		{
			[]int{4, 2, 2, 1},
			2,
		},
		{
			[]int{1},
			1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := majorityElement(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
