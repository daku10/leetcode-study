package main

import (
	"fmt"
	"testing"
)

func TestIsPowerofTwo(t *testing.T) {
	testcases := []struct {
		n    int
		want bool
	}{
		{
			1,
			true,
		},
		{
			16,
			true,
		},
		{
			1073741824,
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPowerOfTwo(tc.n)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
