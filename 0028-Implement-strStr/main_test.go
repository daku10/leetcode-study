package main

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {

	type TestCase struct {
		argHaystack string
		argNeedle string
		expect int
	}

	testCases := []TestCase{
		{"Hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"Hello", "lo", 3},
		{"", "", 0},
		{"", "aa", -1},
		{"aa", "aa", 0},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := strStr(testCase.argHaystack, testCase.argNeedle)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}
