package main

func search(nums []int, target int) int {
	length := len(nums)
    return innerSearch(nums, 0, length - 1, target)
}

func innerSearch(nums []int, l int, r int, target int) int {
	m := (l + r) / 2
	left := nums[l]
	right := nums[r]
	middle := nums[m]
	if left == target {
		return l
	}
	if right == target {
		return r
	}
	if middle == target {
		return m
	}
	if l == r || l == m {
		return -1
	}
	// this looks so stupid... and I'm not sure this is O(log n)?
	can1 := innerSearch(nums, m, r, target)
	can2 := innerSearch(nums, l, m, target)
	if can1 == -1 && can2 == -1 {
		return -1
	} else if can1 != -1 {
		return can1
	} else {
		return can2
	}
}
