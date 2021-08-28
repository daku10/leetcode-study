package main

import (
	"fmt"
	"testing"
)

func TestIntegerToRoman(t *testing.T) {

	type TestCase struct {
		arg int
		expect string
	}

	testCases := []TestCase{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func (t *testing.T)  {
			actual := intToRoman(testCase.arg)
			if actual != testCase.expect {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}
