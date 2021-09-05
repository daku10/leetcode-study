package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {

	type TestCase struct {
		arg int
		expect []string
	}

	testCases := []TestCase{
		{4, []string{"(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"}},
		{3, []string{"((()))","(()())","(())()","()(())","()()()"}},
		{2, []string{"(())", "()()"}},
		{1, []string{"()"}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := generateParenthesis(testCase.arg)
			sort.StringSlice(actual).Sort()
			sort.StringSlice(testCase.expect).Sort()
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("\nactual: %v\nexpect: %v", actual, testCase.expect)
			}
		})
	}
}
