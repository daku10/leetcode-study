package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {

	testcases := []struct {
		arg  string
		want [][]string
	}{
		{
			"aab",
			[][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			"a",
			[][]string{{"a"}},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := partition(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v arg: %v", got, tc.want, tc.arg)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {

	testcases := []struct {
		arg  string
		want bool
	}{
		{
			"a",
			true,
		},
		{
			"ab",
			false,
		},
		{
			"aba",
			true,
		},
		{
			"aa",
			true,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := isPalindrome(tc.arg)
			if tc.want != actual {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
