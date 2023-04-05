package main

import (
	"fmt"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	testcases := []struct {
		nums []int
		want string
	}{
		{
			[]int{10, 2},
			"210",
		},
		{
			[]int{3, 30, 34, 5, 9},
			"9534330",
		},
		{
			[]int{34323, 3432},
			"343234323",
		},
		{
			[]int{432, 43243},
			"43243432",
		},
		{
			[]int{0, 0},
			"0",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := largestNumber(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
