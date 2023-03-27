package main

import (
	"fmt"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	testcases := []struct {
		arg1 string
		arg2 string
		want int
	}{
		{
			"1.01",
			"1.001",
			0,
		},
		{
			"1.0",
			"1.0.0",
			0,
		},
		{
			"1.1.0.2",
			"1.1.2",
			-1,
		},
		{
			"0.1",
			"1.1",
			-1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := compareVersion(tc.arg1, tc.arg2)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
