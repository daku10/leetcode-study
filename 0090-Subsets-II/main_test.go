package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	testcases := []struct {
		arg  []int
		want [][]int
	}{
		{
			[]int{1, 2, 2},
			[][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			[]int{0},
			[][]int{{}, {0}},
		},
		{
			[]int{4, 4, 4, 1, 4},
			[][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := subsetsWithDup(tc.arg)
			if !reflect.DeepEqual(tc.want, actual) {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
