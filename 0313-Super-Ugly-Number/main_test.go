package main

import (
	"fmt"
	"testing"
)

func TestNthSuperUglyNumber(t *testing.T) {
	testcases := []struct {
		n      int
		primes []int
		want   int
	}{
		{
			12,
			[]int{2, 7, 13, 19},
			32,
		},
		{
			1,
			[]int{2, 3, 5},
			1,
		},
		{
			100000,
			[]int{7, 19, 29, 37, 41, 47, 53, 59, 61, 79, 83, 89, 101, 103, 109, 127, 131, 137, 139, 157, 167, 179, 181, 199, 211, 229, 233, 239, 241, 251},
			1092889481,
		},
		{
			15,
			[]int{3, 5, 7, 11, 19, 23, 29, 41, 43, 47},
			35,
		},
		{
			45,
			[]int{2, 3, 7, 13, 17, 23, 31, 41, 43, 47},
			82,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := nthSuperUglyNumber(tc.n, tc.primes)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
