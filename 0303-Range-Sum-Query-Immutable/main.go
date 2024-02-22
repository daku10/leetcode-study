package main

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sums[i] = sum
	}
	return NumArray{nums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.nums[right]
	}
	return this.nums[right] - this.nums[left-1]
}
