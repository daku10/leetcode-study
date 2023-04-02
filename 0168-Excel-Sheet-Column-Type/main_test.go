package main

import (
	"fmt"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	testcases := []struct {
		arg  int
		want string
	}{
		{
			1,
			"A",
		},
		{
			26,
			"Z",
		},
		{
			28,
			"AB",
		},
		{
			52,
			"AZ",
		},
		{
			53,
			"BA",
		},
		{
			701,
			"ZY",
		},
		{
			702,
			"ZZ",
		},
		{
			703,
			"AAA",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := convertToTitle(tc.arg)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
