package main

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	res = append(res, []int{})
	for i := 0; i < length; i++ {
		res = append(res, combination(nums, i+1)...)
	}
	return res
}

func combination(nums []int, k int) [][]int {
	c := make([][]int, 0)
	length := len(nums)
	if k == 1 {
		for i := 0; i < length; i++ {
			c = append(c, []int{nums[i]})
		}
		return c
	}
	if length == k {
		r := make([]int, length)
		copy(r, nums)
		c = append(c, r)
		return c
	}

	// nCk = n * n-1Ck-1 + n-1Ck
	n := nums[0]
	subNums := nums[1:length]
	nCk1 := combination(subNums, k-1)
	nCk1Length := len(nCk1)
	for i := 0; i < nCk1Length; i++ {
		a := nCk1[i]
		a = append(a, n)
		c = append(c, a)
	}
	c = append(c, combination(subNums, k)...)
	return c
}
