package main

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testcases := []struct {
		s    string
		t    string
		want bool
	}{
		{
			"anagram",
			"nagaram",
			true,
		},
		{
			"rat",
			"cat",
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
