package main

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	testcases := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			"a",
			"b",
			false,
		},
		{
			"aa",
			"ab",
			false,
		},
		{
			"aa",
			"aab",
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := canConstruct(tc.ransomNote, tc.magazine)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
