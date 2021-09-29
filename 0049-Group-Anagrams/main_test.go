package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	type TestCase struct {
		arg []string
		expect [][]string
	}

	testCases := []TestCase{
		{[]string{"eat","tea","tan","ate","nat","bat"}, [][]string{{"bat"},{"nat","tan"},{"ate","eat","tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := groupAnagrams(testCase.arg)
			if !reflect.DeepEqual(actual, testCase.expect) {
				t.Errorf("arg: %v, expect: %v, actual: %v", testCase.arg, testCase.expect, actual)
			}
		})
	}
}
