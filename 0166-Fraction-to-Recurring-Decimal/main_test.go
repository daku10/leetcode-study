package main

import (
	"fmt"
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	testcases := []struct {
		numerator   int
		denominator int
		want        string
	}{
		{
			1,
			2,
			"0.5",
		},
		{
			2,
			1,
			"2",
		},
		{
			4,
			333,
			"0.(012)",
		},
		{
			5,
			18,
			"0.2(7)",
		},
		{
			-50,
			8,
			"-6.25",
		},
		{
			-22,
			-2,
			"11",
		},
		{
			-2147483648,
			1,
			"-2147483648",
		},
		{
			7,
			-12,
			"-0.58(3)",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := fractionToDecimal(tc.numerator, tc.denominator)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
