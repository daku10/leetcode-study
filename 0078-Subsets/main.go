package main

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	length := len(nums)
	for i := 0; i < length; i++ {
		resLength := len(res)
		n := nums[i]
		for j := 0; j < resLength; j++ {
			a := res[j]
			tmp := make([]int, len(a))
			copy(tmp, a)
			tmp = append(tmp, n)
			res = append(res, tmp)
		}
	}
	return res
}
