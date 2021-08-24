package main

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	type TestCase struct {
		arg []string
		expect string
	}

	testCases := []TestCase{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"a"}, "a"},
		{[]string{"ab", "a"}, "a"},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			answer := longestCommonPrefix(testCase.arg)
			if answer != testCase.expect {
				t.Errorf("actual: %v, expect: %v", answer, testCase.expect)
			}
		})
	}
}
