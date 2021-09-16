package main

import (
	"fmt"
	"testing"
)

func TestCountAndSay(t *testing.T) {

	type TestCase struct {
		arg int
		expect string
	}

	testCases := []TestCase{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := countAndSay(testCase.arg)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}
