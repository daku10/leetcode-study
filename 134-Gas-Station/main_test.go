package main

import (
	"fmt"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {

	testcases := []struct {
		gas  []int
		cost []int
		want int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
		{
			[]int{3, 1, 1},
			[]int{1, 2, 2},
			0,
		},
		{
			[]int{2},
			[]int{2},
			0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := canCompleteCircuit(tc.gas, tc.cost)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
