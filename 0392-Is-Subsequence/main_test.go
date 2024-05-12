package main

import (
	"fmt"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	testcases := []struct {
		s    string
		t    string
		want bool
	}{
		{
			"abc",
			"ahbgdc",
			true,
		},
		{
			"axc",
			"ahbgdc",
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isSubsequence(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
