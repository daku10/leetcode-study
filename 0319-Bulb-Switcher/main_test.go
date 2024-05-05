package main

import (
	"fmt"
	"testing"
)

func TestBulbSwitch(t *testing.T) {
	testcases := []struct {
		n    int
		want int
	}{
		{
			0, 0,
		},
		{
			1, 1,
		},
		{
			3, 1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := bulbSwitch(tc.n)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
