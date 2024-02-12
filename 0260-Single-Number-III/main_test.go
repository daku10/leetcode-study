package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testcases := []struct {
		nums []int
		want []int
	}{
		{
			[]int{1, 2, 1, 3, 2, 5},
			[]int{3, 5},
		},
		{
			[]int{-1, 0},
			[]int{-1, 0},
		},
		{
			[]int{0, 1},
			[]int{1, 0},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := singleNumber(tc.nums)
			slices.Sort(got)
			slices.Sort(tc.want)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
