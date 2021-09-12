package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
    
	type TestCase struct {
		argNums []int
		argTarget int
		expect int
	}

	testCases := []TestCase{
		{[]int{4,5,6,7,0,1,2}, 0, 4},
		{[]int{4,5,6,7,0,1,2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{4,5,6,7,0,1,2}, 1, 5},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := search(testCase.argNums, testCase.argTarget)
			if actual != testCase.expect {			
				t.Errorf("nums: %v, target: %v, actual: %v, expect: %v", testCase.argNums, testCase.argTarget, actual, testCase.expect)
			}
		})
	}
}
