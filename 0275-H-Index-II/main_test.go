package main

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {
	testcases := []struct {
		citations []int
		want      int
	}{
		{
			[]int{1, 3, 5, 6},
			3,
		},
		{
			[]int{1, 2, 100},
			2,
		},
		{
			[]int{0, 0, 0, 0},
			0,
		},
		{
			[]int{0, 0, 0, 0, 1},
			1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := hIndex(tc.citations)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
