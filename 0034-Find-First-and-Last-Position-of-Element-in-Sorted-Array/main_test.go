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
		{[]int{5,7,7,8,8,10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{0,1,5,8,8,8,8,8,8,8}, 8, []int{3, 9}},
		{[]int{1}, 1, []int{0, 0}},
		{[]int{1,1}, 1, []int{0, 1}},
		{[]int{1, 2, 3}, 1, []int{0, 0}},
		{[]int{1,2,3,3,3,3,4,5,9,}, 3, []int{2, 5}},
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
