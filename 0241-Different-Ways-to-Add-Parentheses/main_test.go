package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {
	testcases := []struct {
		arg  string
		want []int
	}{
		{
			"2-1-1",
			[]int{0, 2},
		},
		{
			"2*3-4*5",
			[]int{-34, -14, -10, -10, 10},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := diffWaysToCompute(tc.arg)
			slices.Sort(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
