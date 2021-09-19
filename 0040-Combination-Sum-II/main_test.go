package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {

	type TestCase struct {
		argCandidates []int
		argTarget     int
		expect        [][]int
	}

	testCases := []TestCase{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := combinationSum2(testCase.argCandidates, testCase.argTarget)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("candidates: %v target: %v expect: %v actual :%v", testCase.argCandidates, testCase.argTarget, testCase.expect, actual)
			}
		})
	}
}
