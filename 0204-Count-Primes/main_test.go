package main

import (
	"fmt"
	"testing"
)

func TestCountPrime(t *testing.T) {
	testcases := []struct {
		n    int
		want int
	}{
		{
			10,
			4,
		},
		{
			0,
			0,
		},
		{
			1,
			0,
		},
		{
			2,
			0,
		},
		{
			3,
			1,
		},
		{
			5000000,
			348513,
		},
		{
			5,
			2,
		},
		{
			6,
			3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countPrimes(tc.n)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
