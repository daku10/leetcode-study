package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T)  {
    
	type TestCase struct {
		arg []int
		expect []int
	}

	testCases := []TestCase{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{1, 2, 4, 3}, []int{1, 3, 2, 4}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{2, 3, 1}, []int{3, 1, 2}},
		{[]int{2, 1, 3, 5, 4}, []int{2, 1, 4, 3, 5}},
		{[]int{4,2,0,2,3,2,0}, []int{4,2,0,3,0,2,2}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			t.Logf("arg: %v\n", testCase.arg)
			nextPermutation(testCase.arg)
			t.Logf("res: %v\n", testCase.arg)
			if !reflect.DeepEqual(testCase.arg, testCase.expect) {
				t.Errorf("exp: %v", testCase.expect)
			}
		})
	}
}
