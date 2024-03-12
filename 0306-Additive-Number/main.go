package main

import (
	"strconv"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	return isAdditiveNumberInternal(num)
}

func isAdditiveNumberInternal(num string) bool {
	for i := 1; i <= len(num)/2; i++ {
		if num[0] == '0' && i > 1 {
			return false
		}
		first := toInt64(num[:i])
	labelJ:
		for j := i + 1; j <= (i + len(num)/2); j++ {
			if num[i] == '0' && j-i > 1 {
				break labelJ
			}
			second := toInt64(num[i:j])
		labelK:
			for k := j + 1; k <= len(num); k++ {
				if num[j] == '0' && k-j > 1 {
					break labelK
				}
				third := toInt64(num[j:k])
				if first+second == third {
					if k == len(num) {
						return true
					}
					if isAdditiveNumberInternal(num[i:]) {
						return true
					}
				}
			}
		}
	}
	return false
}

func toInt64(s string) int64 {
	r, _ := strconv.ParseInt(s, 10, 0)
	return r
}
