package main

func getRow(rowIndex int) []int {
	var prev []int
	var res []int
	for i := 0; i <= rowIndex; i++ {
		res = make([]int, i+1)
		res[0] = 1
		res[i] = 1
		for j := 1; j < i; j++ {
			res[j] = prev[j-1] + prev[j]
		}
		prev = res
	}
	return res
}
