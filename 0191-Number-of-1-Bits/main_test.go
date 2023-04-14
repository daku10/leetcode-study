package main

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	testcases := []struct {
		num  uint32
		want int
	}{
		{
			11,
			3,
		},
		{
			128,
			1,
		},
		{
			4294967293,
			31,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := hammingWeight(tc.num)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
