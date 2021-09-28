package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	
	type TestCase struct {
		arg []int
		expect [][]int
	}

	testCases := []TestCase{
		{[]int{1,1,2}, [][]int{{1,1,2}, {1,2,1}, {2,1,1}}},
		{[]int{1,2,3}, [][]int{{1,2,3}, {1,3,2}, {2,1,3}, {2,3,1}, {3,1,2}, {3,2,1}}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := permuteUnique(testCase.arg)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("nums: %v, expect: %v, actual: %v", testCase.arg, testCase.expect, actual)
			}
		})
	}
}
