package main

import (
	"fmt"
	"testing"
)

func TestToHex(t *testing.T) {
	testcases := []struct {
		num  int
		want string
	}{
		{
			26,
			"1a",
		},
		{
			-1,
			"ffffffff",
		},
		{
			-2,
			"fffffffe",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := toHex(tc.num)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
