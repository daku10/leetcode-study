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
		{[]int{2}, 1, [][]int{{}}},
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1}, 2, [][]int{{1, 1}}},
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
