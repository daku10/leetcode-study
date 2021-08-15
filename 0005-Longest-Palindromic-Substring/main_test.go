package main

import (
	"fmt"
	"testing"
)

func TestPalindromicString(t *testing.T) {

	type TestCase struct {
		arg string
		expect string
	}

	testCases := []TestCase{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"bb", "bb"},
		{"ccc", "ccc"},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func (t *testing.T)  {
			actual := longestPalindrome(testCase.arg)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}

}