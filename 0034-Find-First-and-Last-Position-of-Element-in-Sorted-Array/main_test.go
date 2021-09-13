package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {

	type TestCase struct {
		argNums []int
		argTarget int
		expect []int
	}

	testCases := []TestCase{
		{[]int{5,7,7,8,8,10}, 8, []int{3, 4}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := searchRange(testCase.argNums, testCase.argTarget)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("nums: %v, actual: %v, expect :%v", testCase.argNums, actual, testCase.expect)
			}
		})
	}
}
