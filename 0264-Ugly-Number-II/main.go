package main

import (
	"slices"
)

func nthUglyNumber(n int) int {
	// 1, 2, 3, 2*2, 5, 2*3, 2*2*2, 3*3, 2*5, 2*2*3,
	// 3*5, 2*2*2*2, 2*3*3, 2*2*5, 2*2*2*3, 5*5, -> 16

	// 3*3*3,   2*3*5,   2*2*2*2*2,     2*2*3*3,   2*2*2*5,   3*3*5, 3*2*2*2*2, 2*5*5,
	// 2*3*3*3, 2*2*3*5, 2*2*2*2*2*2*2, 2*2*2*3*3, 2*2*2*2*5,  3*3*3*3, 2*3*3*5, 2*3*2*2*2*2, 2*2*5*5
	// t := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 27, 30, 32, 36, 40, 45, 48, 50, 54, 60, 64, 72, 80, 81, 90, 96, 100}
	// memo := make([]int, n+5)
	memo := []int{1}
	dup := make(map[int]struct{}, n+4)
	i := 0
	for i <= n+1 {
		m := memo[i]
		m2 := m * 2
		m3 := m * 3
		m5 := m * 5
		index := i
		if _, ok := dup[m2]; !ok {
			index++
			memo = append(memo, m2)
			// memo[index] = m2
			dup[m2] = struct{}{}
		}
		if _, ok := dup[m3]; !ok {
			index++
			// memo[index] = m3
			memo = append(memo, m3)
			dup[m3] = struct{}{}
		}
		if _, ok := dup[m5]; !ok {
			index++
			// memo[index] = m5
			memo = append(memo, m5)
			dup[m5] = struct{}{}
		}
		i++
		slices.Sort(memo)
	}

	return memo[n-1]
}
