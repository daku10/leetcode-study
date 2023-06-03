package main

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		res[0] = 1
		res[i] = 1
		for j := i - 1; j > 0; j-- {
			res[j] = res[j-1] + res[j]
		}
	}
	return res
}
