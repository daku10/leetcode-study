package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {

	type TestCase struct {
		argNum []int
		argTarget int
		expect [][]int
	}

	testCases := []TestCase{
		{[]int{1,0,-1,0,-2,2}, 0, [][]int{{-2,-1,1,2},{-2,0,0,2},{-1,0,0,1}}},
		{[]int{2,2,2,2,2}, 8, [][]int{{2,2,2,2}}},
	};
	
	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := fourSum(testCase.argNum, testCase.argTarget)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("actual: %v expect: %v", actual, testCase.expect)
			}
		})
	}
}
