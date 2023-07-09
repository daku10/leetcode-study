package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	testcases := []struct {
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{
			2,
			[][]int{{1, 0}},
			[]int{0, 1},
		},
		{
			4,
			[][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			[]int{0, 1, 2, 3},
		},
		{
			1,
			[][]int{},
			[]int{0},
		},
		{
			2,
			[][]int{{0, 1}, {1, 0}},
			[]int{},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findOrder(tc.numCourses, tc.prerequisites)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
