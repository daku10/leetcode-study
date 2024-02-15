package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testcases := []struct {
		nums []int
		want []int
	}{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			[]int{0},
			[]int{0},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			moveZeroes(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.want) {
				t.Errorf("got: %v, want: %v", tc.nums, tc.want)
			}
		})
	}
}
