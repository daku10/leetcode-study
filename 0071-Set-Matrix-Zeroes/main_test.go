package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetMatrixZeroes(t *testing.T) {
	testcases := []struct {
		arg  [][]int
		want [][]int
	}{
		{
			arg:  [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			arg:  [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			setZeroes(tc.arg)
			if !reflect.DeepEqual(tc.want, tc.arg) {
				t.Errorf("expected: %v, actual: %v", tc.want, tc.arg)
			}
		})
	}
}
