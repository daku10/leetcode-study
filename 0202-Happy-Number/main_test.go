package main

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	testcases := []struct {
		arg  int
		want bool
	}{
		{
			19,
			true,
		},
		{
			2,
			false,
		},
		{
			1,
			true,
		},
		{
			2147483647,
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isHappy(tc.arg)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
