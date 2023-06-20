package main

import (
	"fmt"
	"testing"
)

func TestRangeBitwseAnd(t *testing.T) {
	testcases := []struct {
		left  int
		right int
		want  int
	}{
		{
			5,
			7,
			4,
		},
		{
			0,
			0,
			0,
		},
		{
			1,
			2147483647,
			0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := rangeBitwiseAnd(tc.left, tc.right)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
