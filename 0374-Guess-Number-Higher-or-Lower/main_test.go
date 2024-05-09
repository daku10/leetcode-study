package main

import (
	"fmt"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	testcases := []struct {
		n    int
		pick int
	}{
		{10, 6},
		{1, 1},
		{2, 1},
		{2147483647, 2147483647},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			answer = tc.pick
			got := guessNumber(tc.n)
			if got != tc.pick {
				t.Errorf("got: %v, want: %v", got, tc.pick)
			}
		})
	}
}
