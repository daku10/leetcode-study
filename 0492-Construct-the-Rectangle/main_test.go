package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstructRectangle(t *testing.T) {
	testcases := []struct {
		area int
		want []int
	}{
		{
			4,
			[]int{2, 2},
		},
		{
			6,
			[]int{3, 2},
		},
		{
			37,
			[]int{37, 1},
		},
		{
			122122,
			[]int{427, 286},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := constructRectangle(tc.area)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
