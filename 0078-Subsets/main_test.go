package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	testcases := []struct {
		arg  []int
		want [][]int
	}{
		{
			arg:  []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			arg:  []int{0},
			want: [][]int{{}, {0}},
		},
		{
			arg:  []int{9, 0, 3, 5, 7},
			want: [][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := subsets(tc.arg)
			for i := 0; i < len(actual); i++ {
				sort.IntSlice.Sort(actual[i])
			}
			for i := 0; i < len(tc.want); i++ {
				sort.IntSlice.Sort(tc.want[i])
			}
			sort.Slice(actual, func(i, j int) bool {
				eleI := actual[i]
				eleJ := actual[j]
				if len(eleI) < len(eleJ) {
					return true
				} else if len(eleI) > len(eleJ) {
					return false
				}
				for i := 0; i < len(eleI); i++ {
					if eleI[i] < eleJ[i] {
						return true
					} else if eleI[i] > eleJ[i] {
						return false
					}
				}
				return false
			})
			sort.Slice(tc.want, func(i, j int) bool {
				eleI := tc.want[i]
				eleJ := tc.want[j]
				if len(eleI) < len(eleJ) {
					return true
				} else if len(eleI) > len(eleJ) {
					return false
				}
				for i := 0; i < len(eleI); i++ {
					if eleI[i] < eleJ[i] {
						return true
					} else if eleI[i] > eleJ[i] {
						return false
					}
				}
				return false
			})
			if !reflect.DeepEqual(tc.want, actual) {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
