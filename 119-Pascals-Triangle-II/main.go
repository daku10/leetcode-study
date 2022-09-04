package main

func getRow(rowIndex int) []int {
	var result [][]int
	for i := 0; i <= rowIndex; i++ {
		res := make([]int, i+1)
		res[0] = 1
		res[i] = 1
		for j := 1; j < i; j++ {
			res[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, res)
	}
	return result[rowIndex]
}
