package main

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {

	type TestCase struct {
		s string
		p string
		expect bool
	}

	testCases := []TestCase{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"", ".*", true},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := isMatch(testCase.s, testCase.p)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect : %v", actual, testCase.expect)
			}
		})
	}
}
