package main

import (
	"fmt"
	"reflect"
	"testing"
)

// func nextGreaterElement(nums1 []int, nums2 []int) []int {

// }

func TestNextGreaterElement(t *testing.T) {
	testcases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			[]int{4, 1, 2},
			[]int{1, 3, 4, 2},
			[]int{-1, 3, -1},
		},
		{
			[]int{2, 4},
			[]int{1, 2, 3, 4},
			[]int{3, -1},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := nextGreaterElement(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
