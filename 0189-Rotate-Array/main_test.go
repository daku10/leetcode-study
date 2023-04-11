package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testcases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
		{
			[]int{1, 2, 3},
			0,
			[]int{1, 2, 3},
		},
		{
			[]int{-1},
			2,
			[]int{-1},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			rotate(tc.nums, tc.k)
			if !reflect.DeepEqual(tc.nums, tc.want) {
				t.Errorf("got: %v want: %v", tc.nums, tc.want)
			}
		})
	}
}
