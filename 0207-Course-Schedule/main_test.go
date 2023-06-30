package main

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	testcases := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			2,
			[][]int{
				{1, 0},
			},
			true,
		},
		{
			2,
			[][]int{
				{1, 0},
				{0, 1},
			},
			false,
		},
		{
			2,
			[][]int{
				{0, 1},
			},
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := canFinish(tc.numCourses, tc.prerequisites)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
