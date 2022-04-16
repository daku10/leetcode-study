package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testcases := []struct {
		argNums1 []int
		argM     int
		argNums2 []int
		argN     int
		want     []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			[]int{1},
			1,
			[]int{},
			0,
			[]int{1},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			merge(tc.argNums1, tc.argM, tc.argNums2, tc.argN)
			if !reflect.DeepEqual(tc.want, tc.argNums1) {
				t.Errorf("expected: %v, actual: %v", tc.want, tc.argNums1)
			}
		})
	}
}
