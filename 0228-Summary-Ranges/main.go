package main

import (
	"strconv"
	"strings"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var result []string
	start := nums[0]
	current := nums[0]
	var b strings.Builder
	for i := 1; i < len(nums); i++ {
		if current+1 != nums[i] {
			if start == current {
				result = append(result, strconv.Itoa(start))
			} else {
				b.WriteString(strconv.Itoa(start))
				b.WriteString("->")
				b.WriteString(strconv.Itoa(current))
				result = append(result, b.String())
				b.Reset()
			}
			start = nums[i]
		}
		current = nums[i]
	}

	if start == current {
		result = append(result, strconv.Itoa(start))
	} else {
		b.WriteString(strconv.Itoa(start))
		b.WriteString("->")
		b.WriteString(strconv.Itoa(current))
		result = append(result, b.String())
	}

	return result
}
