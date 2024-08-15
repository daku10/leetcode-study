package main

import (
	"fmt"
	"testing"
)

func TestConvertToBase7(t *testing.T) {
	testcases := []struct {
		num  int
		want string
	}{
		{
			100,
			"202",
		},
		{
			-7,
			"-10",
		},
		{
			6,
			"6",
		},
		{
			10000000,
			"150666343",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := convertToBase7(tc.num)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
