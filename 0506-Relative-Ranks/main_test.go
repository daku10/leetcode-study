package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {
	testcases := []struct {
		score []int
		want  []string
	}{
		{
			[]int{5, 4, 3, 2, 1},
			[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			[]int{10, 3, 8, 9, 4},
			[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
		{
			[]int{1},
			[]string{"Gold Medal"},
		},
		{
			[]int{1, 2},
			[]string{"Silver Medal", "Gold Medal"},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findRelativeRanks(tc.score)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
