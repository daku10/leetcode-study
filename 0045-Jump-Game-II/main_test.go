package main

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {

	type TestCase struct {
		arg []int
		expect int
	}

	testCases := []TestCase{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1, 1, 1, 1, 1}, 4},
		{[]int{2, 1, 3, 1, 4}, 2},
		{[]int{2, 1, 1, 1, 4}, 3},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := jump(testCase.arg)
			if actual != testCase.expect {
				t.Errorf("nums: %v, expect :%d, actual :%d", testCase.arg, testCase.expect, actual)
			}
		})
	}
}
