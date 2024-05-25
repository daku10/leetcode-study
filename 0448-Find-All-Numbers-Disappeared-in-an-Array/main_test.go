package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	testcases := []struct {
		nums []int
		want []int
	}{
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
		{
			[]int{1, 1},
			[]int{2},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findDisappearedNumbers(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
