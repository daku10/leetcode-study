package main

import (
	"fmt"
	"math"
	"testing"
)

func TestArrangeCoins(t *testing.T) {
	testcases := []struct {
		n    int
		want int
	}{
		{
			5,
			2,
		},
		{
			8,
			3,
		},
		{
			1,
			1,
		},
		{
			3,
			2,
		},
		{
			math.MaxInt32,
			65535,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := arrangeCoins(tc.n)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
