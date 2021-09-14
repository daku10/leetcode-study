package main

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {

	type TestCase struct {
		argNums []int
		argTarget int
		expect int
	}

	testCases := []TestCase{
		{[]int{1,3,5,6}, 5, 2},
		{[]int{1,3,5,6}, 2, 1},
		{[]int{1,3,5,6}, 7, 4},
		{[]int{1,3,5,6}, 0, 0},
		{[]int{1}, 0, 0},		
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := searchInsert(testCase.argNums, testCase.argTarget)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v, target: %v, nums: %v", actual, testCase.expect, testCase.argTarget, testCase.argNums)
			}
		})
	}
}
