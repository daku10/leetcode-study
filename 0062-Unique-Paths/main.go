package main

import (
	"math/big"
)

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	bottomCount := m - 1
	leftCount := n - 1
	small := bottomCount
	if small > leftCount {
		small = leftCount
	}

	sum := bottomCount + leftCount
	var result = big.NewInt(1)
	for i := sum; i > sum-small; i-- {
		result.Mul(result, big.NewInt(int64(i)))
	}
	var child = big.NewInt(1)
	for i := 1; i <= small; i++ {
		child.Mul(child, big.NewInt(int64(i)))
	}
	return int(result.Div(result, child).Int64())
}
