package main

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	testcases := []struct {
		columnTitle string
		want        int
	}{
		{
			"A",
			1,
		},
		{
			"Z",
			26,
		},
		{
			"AA",
			27,
		},
		{
			"ZY",
			701,
		},
		{
			"ZZ",
			702,
		},
		{
			"AAA",
			703,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := titleToNumber(tc.columnTitle)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
