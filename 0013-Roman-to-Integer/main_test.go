package main

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {

	type TestCase struct {
		arg string
		expect int
	}

	testCases := []TestCase{
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := romanToInt(testCase.arg)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}
