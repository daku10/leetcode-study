package main

import (
	"fmt"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	testcases := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 3, 4, 2, 2},
			2,
		},
		{
			[]int{3, 1, 3, 4, 2},
			3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findDuplicate(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
