package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
		{
			[]int{2, 3, 4},
			6,
			[]int{1, 3},
		},
		{
			[]int{-1, 0},
			-1,
			[]int{1, 2},
		},
		{
			[]int{0, 0, 3, 4},
			0,
			[]int{1, 2},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := twoSum(tc.numbers, tc.target)
			if !reflect.DeepEqual(actual, tc.want) {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
