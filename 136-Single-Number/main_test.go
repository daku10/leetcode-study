package main

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {

	testcases := []struct {
		arg  []int
		want int
	}{
		{
			[]int{2, 2, 1},
			1,
		},
		{
			[]int{4, 1, 2, 1, 2},
			4,
		},
		{
			[]int{1},
			1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := singleNumber(tc.arg)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
