package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testcases := []struct {
		n    int
		want []string
	}{
		{
			3,
			[]string{"1", "2", "Fizz"},
		},
		{
			5,
			[]string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			15,
			[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := fizzBuzz(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
