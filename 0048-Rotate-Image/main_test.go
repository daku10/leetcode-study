package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {

	type TestCase struct {
		arg [][]int
		expect [][]int
	}

	testCases := []TestCase{
		{[][]int{{1,2,3}, {4,5,6}, {7,8,9}}, [][]int{{7,4,1}, {8,5,2}, {9,6,3}}},
		{[][]int{{5,1,9,11}, {2,4,8,10}, {13,3,6,7}, {15,14,12,16}}, [][]int{{15,13,2,5}, {14,3,4,1}, {12,6,8,9}, {16,7,10,11}}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			rotate(testCase.arg)
			if !reflect.DeepEqual(testCase.arg, testCase.expect) {
				t.Errorf("arg: %v, expect: %v", testCase.arg, testCase.expect)
			}
		})
	}
}
