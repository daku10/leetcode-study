package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	testcases := []struct {
		arg    [][]int
		arg2   []int
		expect [][]int
	}{
		{
			arg:    [][]int{{1, 3}, {6, 9}},
			arg2:   []int{2, 5},
			expect: [][]int{{1, 5}, {6, 9}},
		},
		{
			arg:    [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			arg2:   []int{4, 8},
			expect: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			arg:    [][]int{{1, 2}, {5, 6}},
			arg2:   []int{3, 4},
			expect: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			arg:    [][]int{{1, 2}},
			arg2:   []int{2, 4},
			expect: [][]int{{1, 4}},
		},
		{
			arg:    [][]int{{1, 2}},
			arg2:   []int{3, 4},
			expect: [][]int{{1, 2}, {3, 4}},
		},
		{
			arg:    [][]int{{3, 4}},
			arg2:   []int{1, 2},
			expect: [][]int{{1, 2}, {3, 4}},
		},
		{
			arg:    [][]int{{3, 4}, {5, 6}},
			arg2:   []int{1, 2},
			expect: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			arg:    [][]int{{1, 2}, {5, 6}},
			arg2:   []int{3, 4},
			expect: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			arg:    [][]int{{2, 3}, {4, 5}, {6, 7}},
			arg2:   []int{1, 8},
			expect: [][]int{{1, 8}},
		},
		{
			arg:    [][]int{{1, 2}, {3, 4}},
			arg2:   []int{5, 6},
			expect: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg, tc.arg2), func(t *testing.T) {
			actual := insert(tc.arg, tc.arg2)
			if !reflect.DeepEqual(tc.expect, actual) {
				t.Errorf("expect: %v, actual: %v", tc.expect, actual)
			}
		})
	}
}
