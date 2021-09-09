package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {

	type TestCase struct {
		argDividend int
		argDivisor int
		expect int
	}

	testCases := []TestCase{
		{10, 3, 3},
		{7, -3, -2},
		{0, 1, 0},
		{1, 1, 1},
		{-2147483648, -1, 2147483647},		
		{2, 3, 0},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := divide(testCase.argDividend, testCase.argDivisor)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}
