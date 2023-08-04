package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	testcases := []struct {
		nums []int
		want []string
	}{
		{
			[]int{0, 1, 2, 4, 5, 7},
			[]string{"0->2", "4->5", "7"},
		},
		{
			[]int{0, 2, 3, 4, 6, 8, 9},
			[]string{"0", "2->4", "6", "8->9"},
		},
		{
			[]int{},
			[]string{},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := summaryRanges(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
