package main

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	testcases := []struct {
		s    string
		want int
	}{
		{
			"leetcode",
			0,
		},
		{
			"loveleetcode",
			2,
		},
		{
			"aabb",
			-1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := firstUniqChar(tc.s)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
