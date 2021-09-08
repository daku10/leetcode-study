package main

func removeElement(nums []int, val int) int {

	result := 0
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[i] == val {
			for j := i; j < length; j++ {
				if nums[j] != val {
					nums[result] = nums[j]
					i = j
					result++
					break;
				}
			}	
		} else {
			nums[result] = nums[i]
			result++
		}
	}

	return result
}
