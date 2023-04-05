package main

import (
	"fmt"
	"sort"
	"strings"
)

func compare(x, y string) bool {
	xy := x + y
	yx := y + x
	for i := 0; i < len(xy); i++ {
		if xy[i] > yx[i] {
			return true
		}
		if xy[i] < yx[i] {
			return false
		}
	}
	return true
}

func largestNumber(nums []int) string {
	var result strings.Builder

	sort.Slice(nums, func(i, j int) bool {
		return compare(fmt.Sprint(nums[i]), fmt.Sprint(nums[j]))
	})

	if nums[0] == 0 {
		return "0"
	}

	for _, n := range nums {
		result.WriteString(fmt.Sprint(n))
	}

	return result.String()
}
