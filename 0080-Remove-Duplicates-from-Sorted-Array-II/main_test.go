package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testcases := []struct {
		arg     []int
		want    int
		wantArr []int
	}{
		{
			arg:     []int{1, 1, 1, 2, 2, 3},
			want:    5,
			wantArr: []int{1, 1, 2, 2, 3, 10000},
		},
		{
			[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			7,
			[]int{0, 0, 1, 1, 2, 3, 3, 10000, 10000},
		},
		{
			[]int{1, 2, 3},
			3,
			[]int{1, 2, 3},
		},
		{
			[]int{1, 1, 1, 1, 2, 3},
			4,
			[]int{1, 1, 2, 3, 10000, 10000},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := removeDuplicates(tc.arg)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
			result := true
			for i := 0; i < actual; i++ {
				if tc.arg[i] != tc.wantArr[i] {
					result = false
					break
				}
			}
			if !result {
				t.Errorf("expected: %v, actual: %v", tc.wantArr, tc.arg)
			}
		})
	}
}
