package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testcases := []struct {
		arg  []int
		want []int
	}{
		{
			arg:  []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			arg:  []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
		{
			arg:  []int{2, 0, 2, 1, 1, 0, 0, 2, 2, 2, 1, 2, 0, 1, 2},
			want: []int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			sortColors(tc.arg)
			if !reflect.DeepEqual(tc.want, tc.arg) {
				t.Errorf("expected: %v, actual: %v", tc.want, tc.arg)
			}
		})
	}
}
