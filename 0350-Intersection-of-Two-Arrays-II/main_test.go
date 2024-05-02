package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestIntersect(t *testing.T) {
	testcases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			[]int{1, 2, 2, 1},
			[]int{2, 2},
			[]int{2, 2},
		},
		{
			[]int{4, 9, 5},
			[]int{9, 4, 9, 8, 4},
			[]int{4, 9},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := intersect(tc.nums1, tc.nums2)
			slices.Sort(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
