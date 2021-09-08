package main

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {

	type TestCase struct {
		argNums []int
		argVal int
		expect int
	}

	testCases := []TestCase {
		{[]int{0,1,2,2,3,0,4,2}, 2, 5},
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{}, 1, 0},
		{[]int{1, 1}, 1, 0},
		{[]int{1, 2}, 3, 2},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			t.Logf("original: %v val: %v\n", testCase.argNums, testCase.argVal)
			actual := removeElement(testCase.argNums, testCase.argVal)
			if testCase.expect != actual {
				t.Errorf("actual: %v, expect: %v ", actual, testCase.expect)
			}
			t.Logf("actual: %v, expect: %v nums :%v\n", actual, testCase.expect, testCase.argNums)
		})
	}
}
