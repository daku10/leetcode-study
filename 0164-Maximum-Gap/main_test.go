package main

import (
	"fmt"
	"testing"
)

func TestMaximumGap(t *testing.T) {
	testcases := []struct {
		num      []int
		expected int
	}{
		{
			[]int{3, 6, 9, 1},
			3,
		},
		{
			[]int{10},
			0,
		},
		{
			[]int{1, 1, 1, 1, 1, 5, 5, 5, 5, 5},
			4,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := maximumGap(tc.num)
			if actual != tc.expected {
				t.Errorf("expected: %d, actual: %d", tc.expected, actual)
			}
		})
	}
}
