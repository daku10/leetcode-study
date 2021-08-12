package main

import (
	"reflect"
	"testing"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	panic("unreachable line")
}

func TestTwoSum(t *testing.T) {
	type TestCase struct {
		title string
		nums []int
		target int
		expect []int
	}
	tests := []TestCase{
		{"test1", []int{1, 2, 4}, 6, []int{1, 2}},
		{"test2", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			actual := twoSum(test.nums, test.target)
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("actual: %v, test: %v", actual, test)
			}
		})

	}
}
