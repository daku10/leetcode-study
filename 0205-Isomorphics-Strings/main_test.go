package main

import (
	"fmt"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	testcases := []struct {
		s    string
		t    string
		want bool
	}{
		{
			"egg",
			"add",
			true,
		},
		{
			"foo",
			"bar",
			false,
		},
		{
			"paper",
			"title",
			true,
		},
		{
			"badc",
			"baba",
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isIsomorphic(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
