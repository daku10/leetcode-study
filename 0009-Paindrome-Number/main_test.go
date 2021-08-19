package main

import (
	"fmt"
	"testing"
)

func TestIsPaindrome(t *testing.T) {

	type TestCase struct {
		arg int
		expect bool
	}

	testCases := []TestCase{
		{121, true},
		{-121, false},
		{10, false},
		{-101, false},
		{0, true},		
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := isPalindrome(testCase.arg)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}
