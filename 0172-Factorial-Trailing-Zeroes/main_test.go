package main

import (
	"fmt"
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	testcases := []struct {
		n    int
		want int
	}{
		{
			3,
			0,
		},
		{
			5,
			1,
		},
		{
			6,
			1,
		},
		{
			9,
			1,
		},
		{
			10,
			2,
		},
		{
			0,
			0,
		},
		{
			20,
			4,
		},
		{
			30,
			7,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := trailingZeroes(tc.n)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
