package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMedium(t *testing.T) {

	type TestCase struct {
		arg1 []int
		arg2 []int
		expect float64
	}

	testCases := []TestCase{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0},
		{[]int{}, []int{1}, 1},
		{[]int{2}, []int{}, 2},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			answer := findMedianSortedArrays(testCase.arg1, testCase.arg2)
			if answer != testCase.expect {
				t.Errorf("actual: %v, exect: %v", answer, testCase.expect)
			}
		})
	}

}

func TestCreateArray(t *testing.T) {
	type TestCase struct {
		arg1 []int
		arg2 []int
		expect []int
	}

	testCases := []TestCase{
		{[]int{1, 3}, []int{2}, []int{1, 2, 3}},
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{[]int{0, 0}, []int{0, 0}, []int{0, 0, 0, 0}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{2}, []int{}, []int{2}},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			answer := createSortedArrays(testCase.arg1, testCase.arg2)
			if !reflect.DeepEqual(answer, testCase.expect) {
				t.Errorf("actual: %v, expect: %v", answer, testCase.expect)
			}
		})
	}
}
