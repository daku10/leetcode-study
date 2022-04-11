package main

func sortColors(nums []int) {
	count0 := 0
	count1 := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			count0++
		} else if nums[i] == 1 {
			count1++
		}
	}

	for i := 0; i < length; i++ {
		if count0 > 0 {
			nums[i] = 0
			count0--
		} else if count1 > 0 {
			nums[i] = 1
			count1--
		} else {
			nums[i] = 2
		}
	}
}
