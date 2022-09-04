package main

func generate(numRows int) [][]int {
	result := [][]int{}
	for i := 1; i <= numRows; i++ {
		res := make([]int, i)
		res[0] = 1
		res[i-1] = 1
		for j := 1; j < i-1; j++ {
			res[j] = result[i-1-1][j-1] + result[i-1-1][j]
		}
		result = append(result, res)
	}
	return result
}
