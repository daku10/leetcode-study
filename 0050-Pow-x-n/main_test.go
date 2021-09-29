package main

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {

	type TestCase struct {
		argX float64
		argN int
		expect float64
	}

	testCases := []TestCase{
		{2.0000, 10, 1024.00000},
		{2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
		{8.88023, 3, 700.28148},
		{34.00515, -3, 3e-05},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := myPow(testCase.argX, testCase.argN)
			if actual != testCase.expect {
				t.Errorf("x: %v n: %v expect: %v actual: %v", testCase.argX, testCase.argN, testCase.expect, actual)
			}
		})
	}
}
