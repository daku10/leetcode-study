package main

import (
	"fmt"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	testcases := []struct {
		n    int
		bad  int
		want int
	}{
		{
			5,
			4,
			4,
		},
		{
			3,
			2,
			2,
		},
		{
			1,
			1,
			1,
		},
		{
			5,
			5,
			5,
		},
		{
			2147483647,
			2147483647,
			2147483647,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			BadVersion = tc.bad
			got := firstBadVersion(tc.n)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
