package main

import (
	"fmt"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	testcases := []struct {
		n    int
		want int
	}{
		{10, 12},
		{1, 1},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := nthUglyNumber(tc.n)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
