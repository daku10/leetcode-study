package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCombination(t *testing.T) {
	testcases := []struct {
		argN int
		argK int
		want [][]int
	}{
		{
			argN: 4,
			argK: 2,
			want: [][]int{
				{2, 4},
				{3, 4},
				{2, 3},
				{1, 2},
				{1, 3},
				{1, 4},
			},
		},
		{
			argN: 1,
			argK: 1,
			want: [][]int{{1}},
		},
		{
			argN: 4,
			argK: 3,
			want: [][]int{{4, 3, 2}, {4, 3, 1}, {4, 2, 1}, {3, 2, 1}},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argN, tc.argK), func(t *testing.T) {
			actual := combine(tc.argN, tc.argK)
			for i := 0; i < len(tc.want); i++ {
				sort.IntSlice.Sort(tc.want[i])
			}
			for i := 0; i < len(actual); i++ {
				sort.IntSlice.Sort(actual[i])
			}
			sort.Slice(tc.want, func(i, j int) bool {
				eleI := tc.want[i]
				eleJ := tc.want[j]
				if eleI[0] < eleJ[0] {
					return true
				} else if eleI[0] > eleJ[0] {
					return false
				}
				return eleI[1] < eleJ[1]
			})
			sort.Slice(actual, func(i, j int) bool {
				eleI := actual[i]
				eleJ := actual[j]
				if eleI[0] < eleJ[0] {
					return true
				} else if eleI[0] > eleJ[0] {
					return false
				}
				return eleI[1] < eleJ[1]
			})

			if !reflect.DeepEqual(tc.want, actual) {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
