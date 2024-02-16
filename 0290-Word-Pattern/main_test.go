package main

import (
	"fmt"
	"testing"
)

func TestWordPattern(t *testing.T) {
	testcases := []struct {
		pattern string
		s       string
		want    bool
	}{
		{
			"abba",
			"dog cat cat dog",
			true,
		},
		{
			"abba",
			"dog cat cat fish",
			false,
		},
		{
			"aaaa",
			"dog cat cat dog",
			false,
		},
		{
			"abba",
			"dog dog dog dog",
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := wordPattern(tc.pattern, tc.s)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
