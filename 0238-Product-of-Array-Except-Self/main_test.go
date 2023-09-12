package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	testcases := []struct {
		nums []int
		want []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := productExceptSelf(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
