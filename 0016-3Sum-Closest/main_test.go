package main

import (
	"fmt"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {

	type TestCase struct {
		argNums []int
		argTarget int
		expect int
	}

	testCases := []TestCase{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := threeSumClosest(testCase.argNums, testCase.argTarget)
			if actual != testCase.expect {
				t.Errorf("actual: %v expect: %v", actual, testCase.expect)
			}
		})
	}
}
