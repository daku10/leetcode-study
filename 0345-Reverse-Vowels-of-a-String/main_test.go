package main

import (
	"fmt"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	testcases := []struct {
		s    string
		want string
	}{
		{
			"hello",
			"holle",
		},
		{
			"leetcode",
			"leotcede",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := reverseVowels(tc.s)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
