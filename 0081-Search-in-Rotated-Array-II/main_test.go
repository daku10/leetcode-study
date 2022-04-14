package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	testcases := []struct {
		argNums   []int
		argTarget int
		want      bool
	}{
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			0,
			true,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			3,
			false,
		},
		{
			[]int{1, 0, 1, 1, 1},
			0,
			true,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argNums, tc.argTarget), func(t *testing.T) {
			actual := search(tc.argNums, tc.argTarget)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
