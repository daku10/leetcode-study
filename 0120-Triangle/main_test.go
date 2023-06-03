package main

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	testcases := []struct {
		arg  [][]int
		want int
	}{
		{
			// [[2],[3,4],[6,5,7],[4,1,8,3]]
			[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			11,
		},
		{
			[][]int{{-10}},
			-10,
		},
		{
			//  -1
			//  3,2
			// -3,1,-1
			//[[-1],[3,2],[-3,1,-1]]
			[][]int{{-1}, {3, 2}, {-3, 1, -1}},
			-1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minimumTotal(tc.arg)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
