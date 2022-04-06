package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testcase := []struct {
		arg  [][]int
		want [][]int
	}{
		{
			arg:  [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			arg:  [][]int{{1, 4}, {4, 5}},
			want: [][]int{{1, 5}},
		},
		{
			arg:  [][]int{{1, 4}, {5, 7}, {4, 5}},
			want: [][]int{{1, 7}},
		},
		{
			arg:  [][]int{{2, 3}, {0, 2}},
			want: [][]int{{0, 3}},
		},
		{
			arg:  [][]int{{1, 2}},
			want: [][]int{{1, 2}},
		},
		{
			arg:  [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			want: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		},
		{
			arg:  [][]int{{1, 10}, {3, 4}, {5, 6}, {7, 8}},
			want: [][]int{{1, 10}},
		},
		{
			arg:  [][]int{{1, 4}, {0, 4}},
			want: [][]int{{0, 4}},
		},
	}

	for _, tc := range testcase {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := merge(tc.arg)
			if !reflect.DeepEqual(tc.want, actual) {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
