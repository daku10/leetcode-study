package main

import (
	"fmt"
	"testing"
)

func TestAddStrings(t *testing.T) {
	testcases := []struct {
		num1 string
		num2 string
		want string
	}{
		{
			"11",
			"123",
			"134",
		},
		{
			"456",
			"77",
			"533",
		},
		{
			"1",
			"999",
			"1000",
		},
		{
			"0",
			"0",
			"0",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := addStrings(tc.num1, tc.num2)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
