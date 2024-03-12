package main

type NumArray struct {
	nums []int
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sums[i] = sum
	}
	return NumArray{
		nums,
		sums,
	}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.nums[index]
	this.nums[index] = val
	for i := index; i < len(this.sums); i++ {
		this.sums[i] += diff
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}
	return this.sums[right] - this.sums[left-1]
}
