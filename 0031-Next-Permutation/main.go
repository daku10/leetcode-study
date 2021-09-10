package main

func nextPermutation(nums []int)  {

	length := len(nums)

	for i := length - 1; i >= 0; i-- {
		target := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] < target {
				for k := i; k >= j && k >= 1; k-- {
					nums[k] =  nums[k - 1]
				}
				nums[j] = target
				return
			}
		}
	}
	
	for i := 0; i < length / 2; i++ {
		tmp1 := nums[i]
		nums[i] = nums[length - i - 1]
		nums[length - 1 - i] = tmp1
	}
}
