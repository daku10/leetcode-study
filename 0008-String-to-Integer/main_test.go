package main

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	type TestCase struct {
		arg string
		expect int
	}

	testCases := []TestCase{
		{"42", 42},
		{"    -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"+32", 32},
		{"10000000000000000", 2147483647},
		{"-", 0},
		{"", 0},
		{"  0000000000012345678", 12345678},
		{"0", 0},
		{"00", 0},
		{"000000000000000000", 0},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := myAtoi(testCase.arg)
			if actual != testCase.expect {
				t.Errorf("actual: %v expect: %v", actual, testCase.expect)
			}
		})
	}
}
