package main

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	testcases := []struct {
		arg  []string
		want int
	}{
		{
			[]string{"2", "1", "+", "3", "*"},
			9,
		},
		{
			[]string{"4", "13", "5", "/", "+"},
			6,
		},
		{
			[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			22,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := evalRPN(tc.arg)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
