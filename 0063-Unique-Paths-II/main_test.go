package main

import (
	"fmt"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	testcases := []struct {
		arg  [][]int
		want int
	}{
		{
			arg:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want: 2,
		},
		{
			arg:  [][]int{{0, 1}, {0, 0}},
			want: 1,
		},
		{
			arg:  [][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 0}},
			want: 0,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := uniquePathsWithObstacles(tc.arg)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
