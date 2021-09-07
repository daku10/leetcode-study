package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	type TestCase struct {
		arg []int
		expect int
		expectNums []int
	}

	testCases := []TestCase{
		{[]int{1, 1, 2}, 2, []int{1, 2, 101}},
		{[]int{0,0,1,1,1,2,2,3,3,4}, 5, []int{0, 1, 2, 3, 4, 101, 101, 101, 101, 101}},
		{[]int{1, 1}, 1, []int{1, 101}},
		{[]int{1, 2}, 2, []int{1, 2}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := removeDuplicates(testCase.arg)
			if actual != testCase.expect || !reflect.DeepEqual(testCase.arg, testCase.expectNums) {
				t.Errorf("actual: %v, expect: %v, actualNums: %v, expectNums: %v", actual, testCase.expect, testCase.arg, testCase.expectNums)
			}
		})
	}
}
