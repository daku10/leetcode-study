package main

import (
	"fmt"
	"testing"
)

func TestNumSquare(t *testing.T) {
	testcases := []struct {
		num  int
		want int
	}{
		{
			12,
			3,
		},
		{
			13,
			2,
		},
		{
			34,
			2,
		},
		{
			43,
			3,
		},
		{
			47,
			4,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := numSquares(tc.num)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
