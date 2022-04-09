package main

import (
	"fmt"
	"testing"
)

func TestMinimumPath(t *testing.T) {
	testcases := []struct {
		arg  [][]int
		want int
	}{
		{
			arg:  [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			want: 7,
		},
		{
			arg:  [][]int{{1, 2, 3}, {4, 5, 6}},
			want: 12,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := minPathSum(tc.arg)
			if actual != tc.want {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
