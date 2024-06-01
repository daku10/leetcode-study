package main

import (
	"fmt"
	"testing"
)

func TestRepeatedSubstringPatterns(t *testing.T) {
	testcases := []struct {
		s    string
		want bool
	}{
		{
			"abab",
			true,
		},
		{
			"aba",
			false,
		},
		{
			"abcabcabcabc",
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := repeatedSubstringPattern(tc.s)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
