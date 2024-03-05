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
			[]int{3, 0, 6, 1, 5},
			3,
		},
		{
			[]int{1, 3, 1},
			1,
		},
		{
			[]int{0, 0, 6, 6, 6},
			3,
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
