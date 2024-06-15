package main

import (
	"fmt"
	"testing"
)

func TestFindComplement(t *testing.T) {
	testcases := []struct {
		num  int
		want int
	}{
		{
			5, 2,
		},
		{
			1, 0,
		},
		{
			1<<31 - 1,
			0,
		},
		{
			2,
			1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findComplement(tc.num)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
