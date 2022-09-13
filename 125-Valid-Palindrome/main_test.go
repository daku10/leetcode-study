package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		arg  string
		want bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			" ",
			true,
		},
		{
			"0P",
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPalindrome(tc.arg)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
