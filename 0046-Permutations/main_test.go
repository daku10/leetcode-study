package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
    
	type TestCase struct {
		arg []int
		expect [][]int
	}

	testCases := []TestCase{
		{[]int{1,2,3}, [][]int{{1,2,3}, {1,3,2}, {2,1,3}, {2,3,1}, {3,1,2}, {3,2,1}}},
		{[]int{0,1}, [][]int{{0,1}, {1,0}}},
		{[]int{1}, [][]int{{1}}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := permute(testCase.arg)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("arg: %v, expect: %v, actual: %v", testCase.arg, testCase.expect, actual)
			}
		})
	}
}
