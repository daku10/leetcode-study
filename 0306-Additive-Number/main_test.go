package main

import (
	"fmt"
	"testing"
)

func TestIsAdditiveNumber(t *testing.T) {
	testcases := []struct {
		num  string
		want bool
	}{
		{
			"112358",
			true,
		},
		{
			"199100199",
			true,
		},
		{
			"1",
			false,
		},
		{
			"12",
			false,
		},
		{
			"112",
			true,
		},
		{
			"1023",
			false,
		},
		{
			"101",
			true,
		},
		{
			"000",
			true,
		},
		{
			"12122436",
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isAdditiveNumber(tc.num)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
