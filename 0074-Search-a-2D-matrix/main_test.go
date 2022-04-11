package main

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	testcases := []struct {
		argMatrix [][]int
		argTarget int
		want      bool
	}{
		{
			argMatrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			argTarget: 3,
			want:      true,
		},
		{
			argMatrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			argTarget: 13,
			want:      false,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argMatrix, tc.argTarget), func(t *testing.T) {
			actual := searchMatrix(tc.argMatrix, tc.argTarget)
			if tc.want != actual {
				t.Errorf("expected: %v, actual :%v", tc.want, actual)
			}
		})
	}
}
