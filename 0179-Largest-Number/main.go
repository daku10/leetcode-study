package main

import (
	"fmt"
	"sort"
	"strings"
)

func compare(x, y string) bool {
	i := 0
	for {
		if x[i] > y[i] {
			return true
		}
		if x[i] < y[i] {
			return false
		}
		i++
		if i == len(x) {
			if i < len(y) {
				return compare(x, y[i:])
			}
			return true
		}
		if i == len(y) {
			if i < len(x) {
				return compare(x[i:], y)
			}
			return false
		}
	}
}

func largestNumber(nums []int) string {
	allZero := true
	for _, n := range nums {
		if n != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return "0"
	}
	var result strings.Builder

	sort.Slice(nums, func(i, j int) bool {
		v := compare(fmt.Sprint(nums[i]), fmt.Sprint(nums[j]))
		return v
	})

	for _, n := range nums {
		result.WriteString(fmt.Sprint(n))
	}

	return result.String()
}
