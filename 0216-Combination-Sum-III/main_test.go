package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	testcases := []struct {
		k    int
		n    int
		want [][]int
	}{
		{
			3,
			7,
			[][]int{
				{1, 2, 4},
			},
		},
		{
			3,
			9,
			[][]int{
				{1, 2, 6}, {1, 3, 5}, {2, 3, 4},
			},
		},
		{
			4,
			1,
			[][]int{},
		},
		{
			2,
			17,
			[][]int{{8, 9}},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := combinationSum3(tc.k, tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
