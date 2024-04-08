package main

import (
	"fmt"
	"testing"
)

func TestIsPowerOfFour(t *testing.T) {
	testcases := []struct {
		n    int
		want bool
	}{
		{
			16,
			true,
		},
		{
			5,
			false,
		},
		{
			1,
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPowerOfFour(tc.n)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
