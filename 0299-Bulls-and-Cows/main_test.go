package main

import (
	"fmt"
	"testing"
)

func TestGetHints(t *testing.T) {
	testcases := []struct {
		secret string
		guess  string
		want   string
	}{
		{
			"1807",
			"7810",
			"1A3B",
		},
		{
			"1123",
			"0111",
			"1A1B",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := getHint(tc.secret, tc.guess)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
