package main

import (
	"fmt"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	testcases := []struct {
		g    []int
		s    []int
		want int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2},
			2,
		},
		{
			[]int{1},
			[]int{},
			0,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{10, 9, 8, 7},
			[]int{5, 6, 7, 8},
			2,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findContentChildren(tc.g, tc.s)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
