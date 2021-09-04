package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {

	type TestCase struct {
		arg int
		expect []string
	}

	testCases := []TestCase{
		{3, []string{"((()))","(()())","(())()","()(())","()()()"}},
		{2, []string{"(())", "()()"}},
		{1, []string{"()"}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := generateParenthesis(testCase.arg)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}
