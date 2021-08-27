package main

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {

	type TestCase struct {
		arg []int
		expect int
	}

	testCases := []TestCase{
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{0, 0}, 0},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func (t *testing.T)  {
			actual := maxArea(testCase.arg)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}