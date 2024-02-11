package main

import (
	"fmt"
	"testing"
)

func TestIsUgly(t *testing.T) {
	testcases := []struct {
		n    int
		want bool
	}{
		{
			6,
			true,
		},
		{
			1,
			true,
		},
		{
			14,
			false,
		},
		{
			-1,
			false,
		},
		{
			0,
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isUgly(tc.n)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
