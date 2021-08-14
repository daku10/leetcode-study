package main

import "testing"

func TestLength(t *testing.T) {
	type TestCase struct {
		arg string
		expect int
	}
	testCases := []TestCase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
	}
	for _, test := range testCases {
		t.Run(test.arg, func(t *testing.T) {
			answer := lengthOfLongestSubstring(test.arg)
			if answer != test.expect {
				t.Errorf("actual: %v", answer)
			}
		})		
	}
}
 