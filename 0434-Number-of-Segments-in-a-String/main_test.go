package main

import (
	"fmt"
	"testing"
)

func TestCountSegments(t *testing.T) {
	testcases := []struct {
		s    string
		want int
	}{
		{
			"Hello, my name is John",
			5,
		},
		{
			"Hello",
			1,
		},
		{
			"",
			0,
		},
		{
			"                ",
			0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countSegments(tc.s)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
