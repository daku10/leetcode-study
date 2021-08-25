package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {

	type TestCase struct {
		arg string
		expect bool
	}

	testCases := []TestCase{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := isValid(testCase.arg)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}