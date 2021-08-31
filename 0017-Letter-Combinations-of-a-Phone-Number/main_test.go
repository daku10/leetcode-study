package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {

	type TestCase struct {
		arg string
		expect []string
	}

	testCases := []TestCase{
		{"23", []string{"ad","ae","af","bd","be","bf","cd","ce","cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := letterCombinations(testCase.arg)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("actual: %v expect: %v", actual, testCase.expect)
			}
		})
	}
}
