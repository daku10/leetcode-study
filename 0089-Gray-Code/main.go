package main

func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}
	result := make([]int, 1<<n)
	tmp := make([]int, 1<<n)
	tmp[0] = 0
	tmp[1] = 1
	for i := 2; i <= n; i++ {
		length := 1 << i
		for j := 0; j < length/2; j++ {
			result[j] = tmp[j]
			result[length-j-1] = (1 << (i - 1)) + tmp[j]
		}
		if i != n {
			result, tmp = tmp, result
		}
	}
	return result
}
