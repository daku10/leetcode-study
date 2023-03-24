package main

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	testcases := []struct {
		arg  []int
		want int
	}{
		{
			[]int{3, 4, 5, 1, 2},
			1,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
		},
		{
			[]int{11, 13, 15, 17},
			11,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := findMin(tc.arg)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
