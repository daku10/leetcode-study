package main

import (
	"fmt"
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	testcases := []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			16,
		},
		{
			[][]int{{1}},
			4,
		},
		{
			[][]int{{1, 0}},
			4,
		},
		{
			[][]int{
				{1, 1},
				{1, 1},
			},
			8,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := islandPerimeter(tc.grid)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
