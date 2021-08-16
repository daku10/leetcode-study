package main

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	type TestCase struct {
		s string
		numRows int
		expect string
	}

	testCases := []TestCase{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"}, 
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
		{"ABCD", 2, "ACBD"},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := convert(testCase.s, testCase.numRows)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}