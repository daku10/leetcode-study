package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	type TestCase struct {
		arg int
		expect int
	}
	testCases := []TestCase{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{-2147483648, 0},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			answer := reverse(testCase.arg)
			if answer != testCase.expect {
				t.Errorf("actual: %v, expect: %v", answer, testCase.expect)
			}
		})
	}
}
