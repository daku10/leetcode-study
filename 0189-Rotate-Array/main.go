package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	tmp := make([]int, k)
	copy(tmp, nums[len(nums)-k:])
	copy(nums[k:], nums[:len(nums)-k])
	copy(nums[:k], tmp)
}
