package main

import (
	"fmt"
	"testing"
)

func TestLicenseKeyFormatting(t *testing.T) {
	testcases := []struct {
		s    string
		k    int
		want string
	}{
		{
			"5F3Z-2e-9-w",
			4,
			"5F3Z-2E9W",
		},
		{
			"2-5g-3-J",
			2,
			"2-5G-3J",
		},
		{
			"---",
			3,
			"",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := licenseKeyFormatting(tc.s, tc.k)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
