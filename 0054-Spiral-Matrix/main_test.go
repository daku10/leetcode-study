package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	tests := []struct {
		arg  [][]int
		want []int
	}{
		{
			arg:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			arg:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			arg:  [][]int{{1}},
			want: []int{1},
		},
		{
			arg:  [][]int{{1}, {2}},
			want: []int{1, 2},
		},
		{
			arg:  [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			acutal := spiralOrder(tc.arg)
			if !reflect.DeepEqual(tc.want, acutal) {
				t.Errorf("expect: %v, actual: %v\n", tc.want, acutal)
			}
		})
	}
}
