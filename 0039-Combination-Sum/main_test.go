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
		{[]int{}, 0, [][]int{}},
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
