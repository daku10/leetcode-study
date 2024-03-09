package main

import (
	"fmt"
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	testcases := []struct {
		n    int
		want bool
	}{
		{
			27,
			true,
		},
		{
			25,
			false,
		},
		{
			0,
			false,
		},
		{
			1,
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPowerOfThree(tc.n)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
