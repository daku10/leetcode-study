package main

import (
	"fmt"
	"testing"
)

func TestCanWinNim(t *testing.T) {
	testcases := []struct {
		n    int
		want bool
	}{
		{4, false},
		{1, true},
		{2, true},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := canWinNim(tc.n)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
