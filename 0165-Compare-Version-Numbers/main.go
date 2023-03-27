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

	if len(v1Nums) > len(v2Nums) {
		diff := len(v1Nums) - len(v2Nums)
		for i := 0; i < diff; i++ {
			v2Nums = append(v2Nums, 0)
		}
	} else if len(v1Nums) < len(v2Nums) {
		diff := len(v2Nums) - len(v1Nums)
		for i := 0; i < diff; i++ {
			v1Nums = append(v1Nums, 0)
		}
	}

	for i := range v1Nums {
		if v1Nums[i] > v2Nums[i] {
			return 1
		}
		if v2Nums[i] > v1Nums[i] {
			return -1
		}
	}

	return 0
}
