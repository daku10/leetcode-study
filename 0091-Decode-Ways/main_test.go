package main

import (
	"fmt"
	"testing"
)

func TestNumDecode(t *testing.T) {
	testcases := []struct {
		arg  string
		want int
	}{
		{
			"12", 2,
		},
		{
			"226", 3,
		},
		{
			"06", 0,
		},
		{
			"11106", 2,
		},
		{
			"1201234", 3,
		},
		{
			"111111111111111111111111111111111111111111111", 1836311903,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := numDecodings(tc.arg)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
