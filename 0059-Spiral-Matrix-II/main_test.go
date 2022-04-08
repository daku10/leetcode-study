package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	testcases := []struct {
		arg  int
		want [][]int
	}{
		{
			arg:  3,
			want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			arg:  1,
			want: [][]int{{1}},
		},
		{
			arg:  2,
			want: [][]int{{1, 2}, {4, 3}},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := generateMatrix(tc.arg)
			if !reflect.DeepEqual(tc.want, actual) {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
