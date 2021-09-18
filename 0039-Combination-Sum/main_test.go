package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	
	type TestCase struct {
		argCandidates []int
		argTarget int
		expect [][]int
	}

	testCases := []TestCase{
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2},{2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1}, 2, [][]int{{1, 1}}},
		{[]int{2,3,6,7}, 7, [][]int{{7}, {2,2,3}}},
		{[]int{48,22,49,24,26,47,33,40,37,39,31,46,36,43,45,34,28,20,29,25,41,32,23}, 69, [][]int{}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := combinationSum(testCase.argCandidates, testCase.argTarget)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("actual: %#v, expect: %#v", actual, testCase.expect)
			}
		})
	}
}
