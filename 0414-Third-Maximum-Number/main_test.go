package main

import (
	"fmt"
	"testing"
)

func TestThirdMax(t *testing.T) {
	testcases := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 2, 3},
			1,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{2, 2, 3, 1},
			1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := thirdMax(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
