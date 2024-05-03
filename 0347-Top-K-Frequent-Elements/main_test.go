package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	testcases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{1, 2},
			2,
			[]int{1, 2},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := topKFrequent(tc.nums, tc.k)
			slices.Sort(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
