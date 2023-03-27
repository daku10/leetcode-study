package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1Splitted := strings.Split(version1, ".")
	v2Splitted := strings.Split(version2, ".")
	v1Nums := make([]int, len(v1Splitted))
	for i, n := range v1Splitted {
		a, _ := strconv.Atoi(n)
		v1Nums[i] = a
	}
	v2Nums := make([]int, len(v2Splitted))
	for i, n := range v2Splitted {
		a, _ := strconv.Atoi(n)
		v2Nums[i] = a
	}

	maxLen := len(v1Nums)
	if len(v1Nums) > len(v2Nums) {
		diff := len(v1Nums) - len(v2Nums)
		for i := 0; i < diff; i++ {
			v2Nums = append(v2Nums, 0)
		}
	} else if len(v1Nums) < len(v2Nums) {
		maxLen = len(v2Nums)
		diff := len(v2Nums) - len(v1Nums)
		for i := 0; i < diff; i++ {
			v1Nums = append(v1Nums, 0)
		}
	}

	for i := 0; i < maxLen; i++ {
		var v1, v2 int
		if i < len(v1Nums) {
			v1 = v1Nums[i]
		}
		if i < len(v2Nums) {
			v2 = v2Nums[i]
		}
		if v1 > v2 {
			return 1
		}
		if v2 > v1 {
			return -1
		}
	}

	return 0
}
