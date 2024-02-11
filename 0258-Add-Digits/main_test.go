package main

import (
	"fmt"
	"testing"
)

func TestAddDigits(t *testing.T) {
	testcases := []struct {
		num  int
		want int
	}{
		{
			38,
			2,
		},
		{
			0,
			0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := addDigits(tc.num)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
