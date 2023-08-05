package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testcases := []struct {
		nums []int
		want []int
	}{
		{
			[]int{3, 2, 3},
			[]int{3},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := majorityElement(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
