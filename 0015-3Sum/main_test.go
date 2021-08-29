package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {

	type TestCase struct {
		arg []int
		expect [][]int
	}

	testCases := []TestCase{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, -1}}},
		{[]int{}, [][]int{{}, {}}},
		{[]int{0}, [][]int{{}, {}}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := threeSum(testCase.arg)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("actual: %#v, expect: %#v", actual, testCase.expect)				
			}
		})
	}
}
