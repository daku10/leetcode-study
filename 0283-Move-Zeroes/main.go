package main

func moveZeroes(nums []int) {
	j := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if j == -1 {
				j = i + 1
			}
			for j < len(nums) {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
				j++
			}
		}
	}
}
