package main

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	testcases := []struct {
		arg     []int
		wantSet []int
	}{
		{
			[]int{1, 2, 3, 1},
			[]int{2},
		},
		{
			[]int{1, 2, 1, 3, 5, 6, 4},
			[]int{1, 5},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := findPeakElement(tc.arg)
			found := false
			for _, want := range tc.wantSet {
				if actual == want {
					found = true
				}
			}
			if !found {
				t.Errorf("actual: %v want: %v", actual, tc.wantSet)
			}
		})
	}
}
