package main

import (
	"fmt"
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	testcases := []struct {
		num  int
		want bool
	}{
		{
			16,
			true,
		},
		{
			14,
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPerfectSquare(tc.num)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
