package main

import (
	"fmt"
	"math"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	testcases := []struct {
		x    int
		y    int
		want int
	}{
		{
			1,
			4,
			2,
		},
		{
			3,
			1,
			1,
		},
		{
			math.MaxInt32,
			0,
			31,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := hammingDistance(tc.x, tc.y)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
