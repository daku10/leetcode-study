package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	testcases := []struct {
		arg  []int
		want []int
	}{
		{
			arg:  []int{1, 2, 3},
			want: []int{1, 2, 4},
		},
		{
			arg:  []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 2},
		},
		{
			arg:  []int{9},
			want: []int{1, 0},
		},
		{
			arg:  []int{9, 9},
			want: []int{1, 0, 0},
		},
		{
			arg:  []int{1, 9, 9},
			want: []int{2, 0, 0},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := plusOne(tc.arg)
			if !reflect.DeepEqual(tc.want, actual) {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
